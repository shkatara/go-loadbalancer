package main

import (
	"github.com/gin-gonic/gin"
	"trivago.com/shkatara/goLoadBalancer/routes"
)

func main() {
	// Export env var export DOCKER_API_VERSION=1.39
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", routes.GetAllEvents)

	r.Run(":8080")
}
