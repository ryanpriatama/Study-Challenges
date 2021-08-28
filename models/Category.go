package models

type Category struct {
	ID   uint   `json:"id" form:"id" gorm:"not null"`
	Nama string `json:"nama" form:"nama"`
}
