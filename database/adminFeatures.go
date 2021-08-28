package database

import (
	"SC/config"
	"SC/models"
)

func GetAllSoalInSpecifiedCategory(idcategory int) ([]models.Soal, error) {
	var soal []models.Soal
	if err := config.DB.Find(&soal, "category_id=?", idcategory).Error; err != nil {
		return nil, err
	}

	return soal, nil
}

func DeleteOneSoalSpecifiedId(idcategory int) (interface{}, error) {

	var deleteAQuestion models.Soal
	if err := config.DB.Find(&deleteAQuestion, "id=?", idcategory).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&deleteAQuestion).Error; err != nil {
		return nil, err
	}
	return deleteAQuestion, nil
}

func GetOneQuestionById(soalId int) (models.Soal, error) {
	var soal models.Soal
	if err := config.DB.Where("id=?", soalId).First(&soal).Error; err != nil {
		return soal, err
	}
	return soal, nil
}

func EditSoal(soal models.Soal) (models.Soal, error) {
	if err := config.DB.Save(&soal).Error; err != nil {
		return soal, err
	}
	return soal, nil
}
