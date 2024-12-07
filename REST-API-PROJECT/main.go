package main

import "github.com/gin-gonic/gin"

func main() {
	// this configures default http server that has some basic functionalities
	server := gin.Default()

	// listen on port 8080 while dev it will be localhost
	server.Run(":8080")

}
