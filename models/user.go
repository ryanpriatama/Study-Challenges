package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama      string `json:"nama" form:"nama" gorm:"size:50;not null"`
	Email     string `json:"email" form:"email" gorm:"size:50;not null"`
	Password  string `json:"password" form:"password" gorm:"size:50;not null"`
	TotalPoin int    `json:"poin" form:"poin"`
	Rank      string `json:"rank" form:"rank" gorm:"size:50"`
	Role      string `json:"role" form:"role"`
	Token     string `json:"token" form:"token"`
}
