package models

import "gorm.io/gorm"

type Set_soal_detail struct {
	gorm.Model
	Set_soalID   uint
	SoalID       uint
	Status       string `json:"status" form:"status"`
	Poin         int    `json:"poin" form:"poin"`
	Jawaban_user string `json:"jawaban-user" form:"jawaban-user"`
}
