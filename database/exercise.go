package database

import (
	"SC/config"
	"SC/models"
)

func CreateSetSoal(setSoal models.Set_soal) (models.Set_soal, error) {
	if err := config.DB.Save(&setSoal).Error; err != nil {
		return setSoal, err
	}
	return setSoal, nil
}

func RandomId(soalCategory_id, level uint) []models.Soal {
	var soal []models.Soal
	config.DB.Raw("SELECT id FROM soals WHERE kesulitan_id = ? AND category_id = ? AND approval = 'accept' ORDER BY rand() LIMIT 5", level, soalCategory_id).Scan(&soal)
	return soal
}

func InputSetSoalDetail(setSoalDetail models.Set_soal_detail) models.Set_soal_detail {
	config.DB.Save(&setSoalDetail)
	return setSoalDetail
}

func ShowSetSoal(set_soal_id int) models.Set_soal {
	var set_soal models.Set_soal
	config.DB.Where("id=?", set_soal_id).Find(&set_soal)
	return set_soal
}

func ShowActiveSoal(setSoalId int) []models.Soal {
	var soalDetail []models.Set_soal_detail
	var soal []models.Soal
	config.DB.Raw("select soals.id, soal_pertanyaan, pilihan_a, pilihan_b, pilihan_c, pilihan_d from soals inner join set_soal_details on set_soal_details.soal_id = soals.id where set_soal_details.set_soal_id = ?", setSoalId).Scan(&soal).Scan(&soalDetail)
	return soal
}

func PutAnswer(setSoalId int, jawabanUser map[int]string) []models.Set_soal_detail {
	var soalDetail []models.Set_soal_detail
	config.DB.Where("set_soal_id=?", setSoalId).Find(&soalDetail)
	for i := 0; i < 5; i++ {
		soalDetail[i].Jawaban_user = jawabanUser[i+1]
		soalDetail[i].Status = "answered"
	}
	config.DB.Save(&soalDetail)
	return soalDetail
}

func Scoring(setSoalId int) (int, []int) {
	var soalDetail []models.Set_soal_detail
	var set_soal models.Set_soal

	var totalScore int = 0
	var totalBenar int = 0
	var totalSalah int = 0
	var totalTerjawab int = 0
	var SoalId_salah []int

	config.DB.Where("set_soal_id=?", setSoalId).Find(&soalDetail)
	config.DB.Where("id=?", setSoalId).Find(&set_soal)

	for i := 0; i < 5; i++ {
		var soal models.Soal
		config.DB.Where("id=?", soalDetail[i].SoalID).Find(&soal)
		if soal.Jawaban == soalDetail[i].Jawaban_user {
			totalBenar++
			totalTerjawab++
			switch set_soal.KesulitanID {
			case 1:
				soalDetail[i].Poin = 2
			case 2:
				soalDetail[i].Poin = 3
			case 3:
				soalDetail[i].Poin = 4
			}
		} else if soal.Jawaban != soalDetail[i].Jawaban_user && soalDetail[i].Jawaban_user == "pass" {
			soalDetail[i].Poin = 0
		} else if soal.Jawaban != soalDetail[i].Jawaban_user && soalDetail[i].Jawaban_user != "pass" {
			totalSalah++
			totalTerjawab++
			SoalId_salah = append(SoalId_salah, int(soal.ID))
			soalDetail[i].Poin = -1
		}
		totalScore += soalDetail[i].Poin
	}

	set_soal.TotalBenar = totalBenar
	set_soal.TotalSalah = totalSalah
	set_soal.TotalTerjawab = totalTerjawab
	set_soal.Status = "answered"

	config.DB.Save(&set_soal)
	config.DB.Save(&soalDetail)
	return totalScore, SoalId_salah
}

func UpdateUser(userId, totalScore int) models.User {
	var user models.User
	config.DB.Find(&user, "id=?", userId)
	user.TotalPoin += totalScore

	if user.TotalPoin < 21 {
		user.Rank = "Bronze"
	}
	if user.TotalPoin > 20 && user.TotalPoin < 41 {
		user.Rank = "Silver"
	}
	if user.TotalPoin > 40 {
		user.Rank = "Gold"
	}
	config.DB.Save(&user)
	return user
}

func GetSolution(setSoalId int) []models.Soal {
	var soal []models.Soal

	config.DB.Raw("select soals.id, soals.solusi from soals inner join set_soal_details on set_soal_details.soal_id = soals.id where set_soal_details.set_soal_id = ?", setSoalId).Scan(&soal)
	return soal
}
