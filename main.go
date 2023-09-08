package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	developerName := "Gad Iradufasha"
	projectTitle := "Library API"

	router.GET("/", func(c *gin.Context) {
		c.String(200, strings.ToUpper("Welcome to "+projectTitle+" app, made by "+developerName))
	})

	// Running the server on port 8080
	router.Run(":8080")
}
