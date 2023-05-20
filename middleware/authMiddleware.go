package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/library_management/db"
	"github.com/library_management/helper"
	"github.com/library_management/models"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No Token Found"})
		ctx.Abort()
		return
	}

	var tvalid models.Token
	errDB := db.Db.Where("acc_token = ?",token).Find(&tvalid).Error
	if errDB != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errDB.Error()})
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
