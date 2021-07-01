## GO语言 Docker

```go
package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

	// docker info http://192.168.55.238:2376/info
	fmt.Println(cli.Info(context.Background()))

	// docker ping http://192.168.55.238:2376/_ping
	fmt.Println(cli.Ping(context.Background()))

	// docker client version
	fmt.Println(cli.ClientVersion())

	// docker server version
	fmt.Println(cli.ServerVersion(context.Background()))

	// docker pull
	authConfig := types.AuthConfig{
		Username:      "jusene",
		Password:      "zgx",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	out, err := cli.ImagePull(context.Background(), "docker.io/jusene/pause-amd64:3.2", types.ImagePullOptions{RegistryAuth: authStr})

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)

	// docker tag
	cli.ImageTag(context.Background(), "docker.io/jusene/pause-amd64:3.2", "docker.io/jusene/pause:3.2")
}
```