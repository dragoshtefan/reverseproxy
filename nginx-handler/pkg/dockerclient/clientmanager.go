package dockerclient

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	docker "github.com/docker/docker/client"
)

func GetExternalPort(ctx context.Context, containerName string, internalPort int) (int, error) {
	client, err := defaultDockerClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()

	containerId, err := getContainerId(ctx, client, containerName)
	if err != nil {
		return 0, err
	}

	container, err := client.ContainerInspect(context.Background(), containerId)
	if err != nil {
		return 0, err
	}

	for tcpPort, externalBinding := range container.NetworkSettings.Ports {
		if tcpPort.Port() == fmt.Sprintf("%d",internalPort) {
			return strconv.Atoi(externalBinding[0].HostPort)
		}
	}

	return 0, fmt.Errorf("port %d not found in container %s", internalPort, containerName)
}

func getContainerId(ctx context.Context, client *docker.Client, containerName string) (string, error) {

	filters := filters.NewArgs()
	filters.Add("name", containerName)

	containers, err := client.ContainerList(ctx, container.ListOptions{All: true, Filters: filters})
	if err != nil {
		return "", err
	}

	if len(containers) == 0 {
		return "", fmt.Errorf("container %s not found", containerName)
	}

	return containers[0].ID, nil
}

func defaultDockerClient() (*docker.Client, error) {
	client, err := docker.NewClientWithOpts(
		docker.FromEnv,
		docker.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
