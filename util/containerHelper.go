package util

import (
	"context"
	"io"

	dTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/kelseyhightower/envconfig"
)

type ContainerHelperConfig struct {
	DockerHost string `envconfig:"DOCKER_HOST"`
}

type ContainerHelper struct {
	cli *client.Client
}

func NewContainerHelper() *ContainerHelper {
	var cfg ContainerHelperConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	cli, err := client.NewClientWithOpts(client.WithHost(cfg.DockerHost), client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return &ContainerHelper{cli: cli}
}

func (h *ContainerHelper) LogContainer(ctx context.Context, containerName string) (io.ReadCloser, error) {
	return h.cli.ContainerLogs(ctx, containerName, dTypes.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
}
