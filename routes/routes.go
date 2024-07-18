package routes

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"trivago.com/shkatara/goLoadBalancer/controller"
)

func GetAllEvents(c *gin.Context) {
	ipAddressList := controller.ListContainers()
	min := 0
	max := len(ipAddressList)
	targetServer := rand.Intn(max-min) + min
	requestURL := fmt.Sprintf("http://%s", ipAddressList[targetServer])
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	body, _ := io.ReadAll(res.Body)
	c.JSON(200, gin.H{
		"containerIP": requestURL,
		"data":        string(body),
	})
}
