package routes

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"trivago.com/shkatara/goLoadBalancer/controller"
)

func GetAllEvents(c *gin.Context) {
	ipAddressList := controller.ListContainers()
	min := 0
	max := len(ipAddressList)
	targetServer := rand.Intn(max-min) + min
	requestURL := fmt.Sprintf("http://%s", ipAddressList[targetServer])
	c.Redirect(http.StatusMovedPermanently, requestURL)
}
