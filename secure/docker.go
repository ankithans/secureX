package secure

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var colorReset = "\033[0m"

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

func RunApiContainer() {
	fmt.Println(string(colorPurple), "Spawning up the decoy")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "ankithans/securex_app"

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == imageName {
			fmt.Println(string(colorCyan), "container already running")
			return
		}
	}

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"8080": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("8080"): {{HostIP: "127.0.0.1", HostPort: "8080"}}},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
	time.Sleep(1 * time.Second)
}

func StopApiContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "ankithans/securex_app"

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == imageName {
			fmt.Print(string(colorRed), "Stopping idle container ", container.ID[:10], container.Image, "... ")
			if err := cli.ContainerKill(ctx, container.ID, "SIGHUP"); err != nil {
				panic(err)
			}

			fmt.Println(string(colorWhite), "Successfully Stopped")
		}
	}
}
