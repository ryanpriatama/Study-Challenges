package models

import "gorm.io/gorm"

type Set_soal struct {
	gorm.Model
	KesulitanID   uint
	TotalBenar    int    `json:"total-benar" form:"total-benar"`
	TotalSalah    int    `json:"total-salah" form:"total-salah"`
	TotalTerjawab int    `json:"total-terjawab" form:"total-terjawab"`
	Status        string `json:"status" form:"status"`
	CategoryID    uint
	UserID        uint
}
