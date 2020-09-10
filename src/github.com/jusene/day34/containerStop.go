package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func main() {
	cli, err := client.NewClient("tcp://192.168.55.238:2376", "1.39", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if err := cli.ContainerStop(context.Background(), container.ID, nil); err != nil {
			panic(err)
		}
		fmt.Println(container.Names, " Stop Success")
	}
}
