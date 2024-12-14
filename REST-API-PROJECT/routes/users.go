package routes

import (
	"fmt"
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User data to be created not valid"})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created succesufully"})

}
