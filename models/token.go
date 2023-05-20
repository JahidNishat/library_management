package models

import "gorm.io/gorm"

type Token struct{
	gorm.Model
	AccToken string `json:"acc_token"`
}