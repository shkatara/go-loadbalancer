package controller

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListContainers() []string {
	//var containerList = []string{}
	var ipAddressList = []string{}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()
	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: false})

	if err != nil {
		panic(err)
	}
	for _, ctr := range containers {
		inspect, _ := apiClient.ContainerInspect(context.Background(), ctr.ID)
		//fmt.Println(inspect.NetworkSettings.IPAddress)
		ipAddressList = append(ipAddressList, inspect.NetworkSettings.IPAddress)
	}
	return ipAddressList
}
