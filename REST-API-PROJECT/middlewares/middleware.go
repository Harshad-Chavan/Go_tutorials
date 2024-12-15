package middlewares

import (
	"fmt"
	"net/http"

	"example.com/mod/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	//extract token from incoming request
	token := context.Request.Header.Get("Authorization")

	// these means no token was part of the request
	if token == "" {
		// This abortwithstatus will abort the request will not run the request handler after this
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	err, userid := utils.VeritfyToken(token)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized - Invalid token"})
		return
	}

	context.Next()

}
