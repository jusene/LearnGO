package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"os"
)

func main() {
	cli, err := client.NewClient("tcp://192.168.55.238:2376", "1.39", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// docker push
	authConfig := types.AuthConfig{
		Username: "jusene",
		Password: "zgx",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	out, err := cli.ImagePush(context.Background(), "docker.io/jusene/pause-amd64:3.2", types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
}
