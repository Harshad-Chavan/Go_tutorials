package routes

import (
	"example.com/mod/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//fucntion availble due to in same package

	// set path and handlers
	server.GET("/events", getEvents)

	// one way to call the middleware at the start is to add it before the handlers
	//server.POST("/events", middlewares.Authenticate, createEvents)

	//another apporach to use it for multiple routes ate once
	// create a group and add the route required in the group
	authentcated := server.Group("/")
	authentcated.Use(middlewares.Authenticate)
	authentcated.POST("/events", createEvents)
	authentcated.PUT("/events/:id", updateEvent)
	authentcated.DELETE("/events/:id", deleteEvent)

	// to get dynamic ids ,values use ':var_name'
	server.GET("/events/:id", getEvent)

	//User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)

}
