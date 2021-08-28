package database

import (
	"SC/config"
	"SC/models"
)

func CreateQuestion(soal models.Soal) (models.Soal, error) {
	if err := config.DB.Save(&soal).Error; err != nil {
		return soal, err
	}
	return soal, nil
}
