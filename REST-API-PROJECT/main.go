package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// this configures default http server that has some basic functionalities
	server := gin.Default()

	// set path and handlers
	server.GET("/events", getEvents)

	// listen on port 8080 while dev it will be localhost
	server.Run(":8080")

}

// context is passsed automatically by gin to a handle function
func getEvents(context *gin.Context) {
	// map is also valid
	//context.JSON(http.StatusOK, map[string]string{"message": "Hello"})

	// gin .H has same structure as the above map map[string]any
	context.JSON(http.StatusOK, gin.H{"message": "Hello !"})

}
