package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library_management/helper"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == ""{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No Token Found"})
		ctx.Abort()
		return
	}

	claims, err := helper.Validation(token)
	if err != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}

	ctx.Set("email", claims.Email)
	ctx.Set("first_name", claims.FirstName)
	ctx.Set("last_name", claims.LastName)
	ctx.Set("user_type", claims.UserType)
	ctx.Set("user_id", claims.UserId)

	ctx.Next()
}
