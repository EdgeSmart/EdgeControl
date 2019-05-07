package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

const (
	host   = "http://%s.dapi.hopeness.net"
	port   = 2375
	apiVer = "1.39"
)

// New create a new docker client
func New(clusterName string) (context.Context, *client.Client, error) {
	clientContext := context.Background()
	header := map[string]string{}
	dockerClient, err := client.NewClient(fmt.Sprintf(host+":%d", clusterName, port), apiVer, nil, header)
	if err != nil {
		return nil, nil, err
	}
	return clientContext, dockerClient, nil
}
