package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userid := context.GetInt64("userid")

	// comtext param used to fetch values from the url
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not  parse requested id "})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch requested event"})
		return
	}

	err = event.Register(userid)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register requested event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registered for the event"})

}

func cancelForEvent(context *gin.Context) {

}
