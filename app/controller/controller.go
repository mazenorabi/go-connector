package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBiEvents(c *gin.Context) {
	var event models.BiEvent

	if err := c.BindJSON(&event); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, event)
}
