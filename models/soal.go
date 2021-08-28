package models

import "gorm.io/gorm"

type Soal struct {
	gorm.Model
	Soal_pertanyaan string `json:"soal-pertanyaan" form:"soal-pertanyaan" gorm:"not null"`
	PilihanA        string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB        string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC        string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD        string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban         string `json:"jawaban" form:"jawaban"`
	KesulitanID     uint
	Solusi          string `json:"solusi" form:"solusi"`
	Approval        string `json:"approval" form:"approval"`
	CategoryID      uint
}
