package controller

import (
	"SC/database"
	"SC/models"
	"net/http"

	"github.com/labstack/echo"
)

func SubmitQuestionAdmin(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	submitSoal := models.Soal{}
	submitSoal.Approval = "accept"
	c.Bind(&submitSoal)
	soal, err := database.CreateQuestion(submitSoal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new question added",
		"data":    soal,
	})
}

func SubmitQuestion(c echo.Context) error {
	submitSoal := models.Soal{}
	submitSoal.Approval = "not yet"
	c.Bind(&submitSoal)
	soal, err := database.CreateQuestion(submitSoal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapSoal := map[string]interface{}{
		"Soal":      soal.ID,
		"Pilihan A": soal.PilihanA,
		"Pilihan B": soal.PilihanB,
		"Pilihan C": soal.PilihanC,
		"Pilihan D": soal.PilihanD,
		"Jawaban":   soal.Jawaban,
		"Kesulitan": soal.KesulitanID,
		"Solusi":    soal.Solusi,
		"Kategori":  soal.CategoryID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new question added",
		"data":    mapSoal,
	})
}
