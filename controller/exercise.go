package controller

import (
	"SC/database"
	"SC/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func GenerateRandomQuestion(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(userId, c); err != nil {
		return err
	}

	setSoal := models.Set_soal{
		TotalBenar:    0,
		TotalSalah:    0,
		TotalTerjawab: 0,
		Status:        "Not Answered Yet",
		UserID:        uint(userId),
	}
	c.Bind(&setSoal)
	newSetSoal, err := database.CreateSetSoal(setSoal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}

	randomizeSoal(newSetSoal.ID, newSetSoal.CategoryID, newSetSoal.KesulitanID, c)

	materi, level := levelAndCategoryIDConvert(newSetSoal.CategoryID, newSetSoal.KesulitanID)
	mapSetSoal := map[string]interface{}{
		"ID":             newSetSoal.ID,
		"UserID":         newSetSoal.UserID,
		"Mata Pelajaran": materi,
		"Kesulitan":      level,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new set soal added",
		"data":    mapSetSoal,
	})
}

func levelAndCategoryIDConvert(materi, level uint) (string, string) {
	var materiConvert string
	var levelConvert string

	switch materi {
	case 1:
		materiConvert = "Kimia"
	case 2:
		materiConvert = "Fisika"
	case 3:
		materiConvert = "Biologi"
	case 4:
		materiConvert = "Matematika"
	case 5:
		materiConvert = "Bahasa Inggris"
	case 6:
		materiConvert = "Ekonomi"
	case 7:
		materiConvert = "Geografi"
	default:
		materiConvert = "Unknown"
	}

	switch level {
	case 1:
		levelConvert = "Easy"
	case 2:
		levelConvert = "Medium"
	case 3:
		levelConvert = "Hard"
	default:
		levelConvert = "unknown"
	}
	return materiConvert, levelConvert
}

func randomizeSoal(soal_id, soalCategory_id, level uint, c echo.Context) {
	random := database.RandomId(soalCategory_id, level)
	for i := 0; i < 5; i++ {
		newSetSoalDetail := models.Set_soal_detail{
			Set_soalID:   soal_id,
			SoalID:       random[i].ID,
			Status:       "not answered",
			Poin:         0,
			Jawaban_user: "pass",
		}
		database.InputSetSoalDetail(newSetSoalDetail)
	}
}

func ShowActiveQuestion(c echo.Context) error {
	userId, err1 := strconv.Atoi(c.Param("user_id"))
	setSoalId, err2 := strconv.Atoi(c.Param("set_soal_id"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	setSoal := database.ShowSetSoal(setSoalId)
	if setSoal.UserID != uint(userId) || setSoal.Status == "answered" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid set soal id",
		})
	}
	if err := UserAuthorize(userId, c); err != nil {
		return err
	}

	showSoal := database.ShowActiveSoal(setSoalId)
	type arraySoal struct {
		Soal      string
		Soal_id   uint
		Pilihan_A string
		Pilihan_B string
		Pilihan_C string
		Pilihan_D string
	}
	var mapshowSoal []arraySoal

	for i := 0; i < 5; i++ {
		newArray := arraySoal{
			Soal:      showSoal[i].Soal_pertanyaan,
			Soal_id:   showSoal[i].ID,
			Pilihan_A: showSoal[i].PilihanA,
			Pilihan_B: showSoal[i].PilihanB,
			Pilihan_C: showSoal[i].PilihanC,
			Pilihan_D: showSoal[i].PilihanD,
		}
		mapshowSoal = append(mapshowSoal, newArray)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All questions show successfully",
		"data":    mapshowSoal,
	})
}

func AnswerQuestion(c echo.Context) error {
	userId, err1 := strconv.Atoi(c.Param("user_id"))
	setSoalId, err2 := strconv.Atoi(c.Param("set_soal_id"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	setSoal := database.ShowSetSoal(setSoalId)
	if setSoal.UserID != uint(userId) || setSoal.Status == "answered" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid set soal id",
		})
	}
	if err := UserAuthorize(userId, c); err != nil {
		return err
	}

	type jawaban struct {
		Soal_1 string
		Soal_2 string
		Soal_3 string
		Soal_4 string
		Soal_5 string
	}
	var jawabanUser jawaban
	c.Bind(&jawabanUser)

	var jawabanUserMap = make(map[int]string)
	jawabanUserMap[1] = jawabanUser.Soal_1
	jawabanUserMap[2] = jawabanUser.Soal_2
	jawabanUserMap[3] = jawabanUser.Soal_3
	jawabanUserMap[4] = jawabanUser.Soal_4
	jawabanUserMap[5] = jawabanUser.Soal_5
	database.PutAnswer(setSoalId, jawabanUserMap)
	totalScore, SoalId_salah := database.Scoring(setSoalId)

	userUpdated := database.UpdateUser(userId, totalScore)

	//Cache
	cache := Redis()
	key := fmt.Sprintf("userId%d", userId)
	json, _ := json.Marshal(userUpdated)
	cache.Set(ctx, key, json, time.Minute)

	mapResult := map[string]interface{}{
		"Total Score": totalScore,
		"Soal Salah":  SoalId_salah,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Question Answered",
		"data":    mapResult,
	})
}

func ShowSolution(c echo.Context) error {
	userId, err1 := strconv.Atoi(c.Param("user_id"))
	setSoalId, err2 := strconv.Atoi(c.Param("set_soal_id"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	setSoal := database.ShowSetSoal(setSoalId)
	if setSoal.UserID != uint(userId) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid set soal id",
		})
	}
	if err := UserAuthorize(userId, c); err != nil {
		return err
	}

	solution := database.GetSolution(setSoalId)
	mapSolution := map[string]string{
		"Solusi 1": solution[0].Solusi,
		"Solusi 2": solution[1].Solusi,
		"Solusi 3": solution[2].Solusi,
		"Solusi 4": solution[3].Solusi,
		"Solusi 5": solution[4].Solusi,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Here is the solutions",
		"data":    mapSolution,
	})
}
