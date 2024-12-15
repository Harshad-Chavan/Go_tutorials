package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/mod/models"
	"example.com/mod/utils"
	"github.com/gin-gonic/gin"
)

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

	//extract token from incoming request
	token := context.Request.Header.Get("Authorization")

	// these means no token was part of the request
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	err, userid := utils.VeritfyToken(token)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// if some fields are missing int he request content
	// null values will be considered if we want to force then
	// add tags in event struct
	err = context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserID = userid

	err = event.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save  data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}

func getEvent(context *gin.Context) {

	// comtext param used to fetch values from the url
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not  requested id "})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch requested event"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {

	// comtext param used to fetch values from the url
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not  requested id "})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch requested event"})
		return
	}

	var updatedevent models.Event

	err = context.ShouldBindJSON(&updatedevent)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedevent.ID = eventId
	err = updatedevent.Update()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})

}

func deleteEvent(context *gin.Context) {

	// comtext param used to fetch values from the url
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not  requested id "})
		return
	}

	// to check event exists
	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch requested event"})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete requested event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}
