package main

import (
	"example.com/mod/db"

	"github.com/gin-gonic/gin"

	"example.com/mod/routes"
)

func main() {
	//initialzie db
	db.InitDB("sqlite3")

	// this configures default http server that has some basic functionalities
	server := gin.Default()
	routes.RegisterRoutes(server)

	// listen on port 8080 while dev it will be localhost
	server.Run(":8080")

}
