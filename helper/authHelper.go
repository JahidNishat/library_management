package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context) (err error) {
	userType := ctx.GetString("user_type")

	if userType != "ADMIN" {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return nil
}

func CheckUserId(ctx *gin.Context, uid string) (err error) {
	userType := ctx.GetString("user_type")
	userId := ctx.GetString("user_id")

	if userType == "USER" && userId != uid {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return nil
}
