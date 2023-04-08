package controller

import (
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/library_management/config"
	"github.com/library_management/helper"
	"github.com/library_management/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init(){
	DB = config.Connect()
}

func SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Creating Hashing Password
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	//Creating Token
	token, refreshToken, err := helper.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.UserType, user.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hash
	user.Token = token
	user.RefreshToken = refreshToken

	//Save To Database
	if err := DB.Create(&user).Error; err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func LogIn(ctx *gin.Context) {
	var (
		input models.UserLogIn
		user models.User
	)
	
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Check Email ID
	if err := DB.Where("email = ?", input.Email).Find(&user).Error; err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email ID"})
		return
	}

	//Check Password
	ok, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password doesn't match"})
		return
	}

	token, refreshToken, err := helper.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.UserType, user.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Token = token
	user.RefreshToken = refreshToken

	if err := DB.Updates(&user).Error; err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
