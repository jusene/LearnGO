package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
)

func main() {
	cli, err := client.NewClient("tcp://192.168.55.238:2376", "1.39", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	containerConfig := container.Config{
		Tty:       true,
		OpenStdin: true,
		Image:     "nginx",
	}

	hostConfig := container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					"0.0.0.0",
					"80",
				},
			},
		},
	}

	networkConfig := network.NetworkingConfig{}
	containerResp, err := cli.ContainerCreate(context.Background(), &containerConfig, &hostConfig, &networkConfig, "nginx")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(context.Background(), containerResp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
