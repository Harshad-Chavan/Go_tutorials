package main

import (
	"net/http"

	"example.com/mod/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// this configures default http server that has some basic functionalities
	server := gin.Default()

	// set path and handlers
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	// listen on port 8080 while dev it will be localhost
	server.Run(":8080")

}

// context is passsed automatically by gin to a handle function
func getEvents(context *gin.Context) {
	// map is also valid
	//context.JSON(http.StatusOK, map[string]string{"message": "Hello"})

	//use the getenvents func from models
	events := models.GetAllEvents()

	// gin .H has same structure as the above map map[string]any
	//context.JSON(http.StatusOK, gin.H{"message": "Hello !"})

	context.JSON(http.StatusOK, events)

}

func createEvents(context *gin.Context) {
	var event models.Event

	// if some fields are missing int he request content
	// null values will be considered if we want to force then
	// add tags in event struct
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
