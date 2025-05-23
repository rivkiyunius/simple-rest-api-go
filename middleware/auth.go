package middleware

import (
	"event-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token not provided."})
		return
	}
	userId, err := utils.VerifyJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token."})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
