package controller

import (
	"github.com/library_management/db"
	"github.com/library_management/models"
)

func CreateToken(token string){
	var tvalid models.Token
	tvalid.AccToken = token
	db.Db.Save(&tvalid)
}

func DeleteToken(token string)(err error){
	var tvalid models.Token
	errDB := db.Db.Where("acc_token = ?",token).Find(&tvalid).Error
	if errDB != nil{
		return errDB
	}

	db.Db.Delete(&tvalid)

	return nil
}