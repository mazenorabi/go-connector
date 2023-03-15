package main

import (
	"app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Weclome to HTG Events Connector"})
	})

	router.POST("/bievents", controllers.ReadBiEvents)

	router.Run()
}
