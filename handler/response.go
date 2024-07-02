package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorz struct {
	Message string `json:"message"`
}
type responsestat struct {
	Stat string `json:"message"`
}

func errorzResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorz{Message: message})
}
