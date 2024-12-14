package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//fucntion availble due to in same package

	// set path and handlers
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	// to get dynamic ids ,values use ':var_name'
	server.GET("/events/:id", getEvent)
	// Update and event data
	server.PUT("/events/:id", updateEvent)
	// Delte an event
	server.DELETE("/events/:id", deleteEvent)

	//signup route for user
	server.POST("/signup")

}
