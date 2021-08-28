package controller

import (
	"SC/auth"
	"SC/config"
	"SC/database"
	"SC/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//AUTHORIZED AND ROLE AS ADMIN
func AuthorizedAdmin(c echo.Context) bool {
	_, role := auth.ExtractTokenUserId(c)

	if role != "admin" {
		return false
	}
	return true
}

//ADMIN FEATURES: EDIT QUESTION
func EditQuestion(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	soalId, err2 := strconv.Atoi(c.Param("soalId"))
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	oneProblem, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "cannot find soal based on id")
	}
	c.Bind(&oneProblem)
	editedSoal, err := database.EditSoal(oneProblem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot edit soal")
	}
	mapQuestion := map[string]interface{}{
		"ID":              editedSoal.ID,
		"Soal_pertanyaan": editedSoal.Soal_pertanyaan,
		"PilihanA":        editedSoal.PilihanA,
		"PilihanB":        editedSoal.PilihanB,
		"PilihanC":        editedSoal.PilihanC,
		"PilihanD":        editedSoal.PilihanD,
		"Jawaban":         editedSoal.Jawaban,
		"KesulitanID":     editedSoal.KesulitanID,
		"Solusi":          editedSoal.Solusi,
		"Approval":        editedSoal.Approval,
		"CategoryID":      editedSoal.CategoryID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated Problem",
		"Data":    mapQuestion,
	})
}

//ADMIN FEATURES: UPDATING APPROVAL STATUS (ACCEPT, REJECT, AND NOT YET)
func EditSubmitQuestion(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid problem id",
		})
	}

	oneProblem, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "cannot find soal based on id")
	}
	c.Bind(&oneProblem)
	_, err1 := database.EditSoal(oneProblem)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, "Cannot Edit Status Approval")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated Approval",
	})
}

//ADMIN FEATURES: GET ALL QUESTION BASED ON CATEGORY
func GetQuestionByCategory(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	categoryId, err := strconv.Atoi(c.Param("MataPelajaranId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid category id",
		})
	}

	soalByCategoryList, err := database.GetAllSoalInSpecifiedCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find problem based on the category id",
		})
	}

	type arrayOfSoal struct {
		ID              uint
		Soal_pertanyaan string
		Pilihan_A       string
		Pilihan_B       string
		Pilihan_C       string
		Pilihan_D       string
		Jawaban         string
		KesulitanID     uint
		Solusi          string
		Approval        string
		CategoryId      uint
	}
	var mapArraySoal []arrayOfSoal

	for i := 0; i < len(soalByCategoryList); i++ {
		newArray := arrayOfSoal{
			ID:              soalByCategoryList[i].ID,
			Soal_pertanyaan: soalByCategoryList[i].Soal_pertanyaan,
			Pilihan_A:       soalByCategoryList[i].PilihanA,
			Pilihan_B:       soalByCategoryList[i].PilihanB,
			Pilihan_C:       soalByCategoryList[i].PilihanC,
			Pilihan_D:       soalByCategoryList[i].PilihanD,
			Jawaban:         soalByCategoryList[i].Jawaban,
			KesulitanID:     soalByCategoryList[i].KesulitanID,
			Solusi:          soalByCategoryList[i].Solusi,
			Approval:        soalByCategoryList[i].Approval,
			CategoryId:      soalByCategoryList[i].CategoryID,
		}

		mapArraySoal = append(mapArraySoal, newArray)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapArraySoal,
	})

}

//ADMIN FEATURES: GET QUESTION BASED ON ID
func GetQuestionById(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid problem id",
		})
	}
	soal, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot find the problem",
		})
	}
	mapQuestion := map[string]interface{}{
		"ID":              soal.ID,
		"Soal_pertanyaan": soal.Soal_pertanyaan,
		"PilihanA":        soal.PilihanA,
		"PilihanB":        soal.PilihanB,
		"PilihanC":        soal.PilihanC,
		"PilihanD":        soal.PilihanD,
		"Jawaban":         soal.Jawaban,
		"KesulitanID":     soal.KesulitanID,
		"Solusi":          soal.Solusi,
		"Approval":        soal.Approval,
		"CategoryID":      soal.CategoryID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapQuestion,
	})
}

//ADMIN FEATURES: DELETE QUESTION BASED ON ID
func DeleteQuestion(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	_, err1 := database.DeleteOneSoalSpecifiedId(soalId)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete soal",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Deleted Problem",
	})

}

//ADMIN FEATURES: SHOW ALL PROBLEM THAT HAS NOT BEEN REVIEWED -- BY CATEGORY
func ShowSubmittedQuestion(c echo.Context) error {
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	categoryId, err := strconv.Atoi(c.Param("kategori_materi_pelajaran_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot use the id",
		})
	}

	var soal []models.Soal
	if err := config.DB.Where(map[string]interface{}{"category_id": categoryId, "approval": "not yet"}).Find(&soal).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot find the problem that needs approval")
	}

	type arrayOfSoal struct {
		ID              uint
		Soal_pertanyaan string
		Pilihan_A       string
		Pilihan_B       string
		Pilihan_C       string
		Pilihan_D       string
		Jawaban         string
		KesulitanID     uint
		Solusi          string
		Approval        string
		CategoryId      uint
	}
	var mapArraySoal []arrayOfSoal

	for i := 0; i < len(soal); i++ {
		newArray := arrayOfSoal{
			ID:              soal[i].ID,
			Soal_pertanyaan: soal[i].Soal_pertanyaan,
			Pilihan_A:       soal[i].PilihanA,
			Pilihan_B:       soal[i].PilihanB,
			Pilihan_C:       soal[i].PilihanC,
			Pilihan_D:       soal[i].PilihanD,
			Jawaban:         soal[i].Jawaban,
			KesulitanID:     soal[i].KesulitanID,
			Solusi:          soal[i].Solusi,
			Approval:        soal[i].Approval,
			CategoryId:      soal[i].CategoryID,
		}

		mapArraySoal = append(mapArraySoal, newArray)

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapArraySoal,
	})
}
