package main

import (
	"context"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClient("tcp://192.168.55.238:2376", "1.39", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	cli.ImageBuild(context.Background())
}
