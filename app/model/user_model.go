package model

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Mail_address string `json:"mailAddress"`
	User_password string `json:"password"`
}

type User struct {
    gorm.Model
    Id  string `gorm:"primaryKey" json:"id"`
    Mail_address string `json:"mail_address"`
	Password string `json:"password"`
}