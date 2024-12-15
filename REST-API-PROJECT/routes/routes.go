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

	// Events routess
	authentcated.Use(middlewares.Authenticate)
	authentcated.POST("/events", createEvents)
	authentcated.PUT("/events/:id", updateEvent)
	authentcated.DELETE("/events/:id", deleteEvent)
	// to get dynamic ids ,values use ':var_name'
	server.GET("/events/:id", getEvent)

	//registration routes
	authentcated.POST("/events/:id/register", registerForEvent)
	authentcated.DELETE("/events/:id/register", cancelForEvent)

	//User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)

}
