package main

import (
	"fmt"
	"net/http"

	"example.com/mod/db"
	"example.com/mod/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialzie db
	db.InitDB("sqlite3")

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
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch requested data"})
		return
	}

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
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserID = 1
	err = event.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save  data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
