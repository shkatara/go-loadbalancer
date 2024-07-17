package main

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func listContainers() []string {
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

func GetAllEvents(c *gin.Context) {
	ipAddressList := listContainers()
	c.JSON(200, gin.H{
		"ContainersList": ipAddressList,
	})
}

func main() {
	// Export env var export DOCKER_API_VERSION=1.39
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", GetAllEvents)

	r.Run(":8080")
}
