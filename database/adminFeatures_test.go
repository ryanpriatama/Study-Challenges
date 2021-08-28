package database

import (
	"SC/config"
	"SC/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBSoal = models.Soal{
		Soal_pertanyaan: "1 + 9 = ?",
		PilihanA:        "10",
		PilihanB:        "11",
		PilihanC:        "12",
		PilihanD:        "19",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "a",
		Approval:        "accept",
		CategoryID:      2,
	}

	mockDBSoalEdit = models.Soal{
		Soal_pertanyaan: "99 + 2 = ?",
		PilihanA:        "100",
		PilihanB:        "110",
		PilihanC:        "120",
		PilihanD:        "101",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "d",
		Approval:        "accept",
		CategoryID:      2,
	}

	mockDBSoal1 = models.Soal{
		Soal_pertanyaan: "109 + 2 = ?",
		PilihanA:        "100",
		PilihanB:        "111",
		PilihanC:        "120",
		PilihanD:        "101",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "b",
		Approval:        "accept",
		CategoryID:      2,
	}

	mockDBSoal2 = models.Soal{
		Soal_pertanyaan: "99 + 2 = ?",
		PilihanA:        "100",
		PilihanB:        "110",
		PilihanC:        "120",
		PilihanD:        "101",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "d",
		Approval:        "accept",
		CategoryID:      2,
	}

	mockDBSoal3 = models.Soal{
		Soal_pertanyaan: "199 + 2 = ?",
		PilihanA:        "100",
		PilihanB:        "310",
		PilihanC:        "220",
		PilihanD:        "201",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "d",
		Approval:        "accept",
		CategoryID:      1,
	}

	mockDBSoal4 = models.Soal{
		Soal_pertanyaan: "299 + 2 = ?",
		PilihanA:        "200",
		PilihanB:        "410",
		PilihanC:        "301",
		PilihanD:        "220",
		Jawaban:         "",
		KesulitanID:     1,
		Solusi:          "c",
		Approval:        "accept",
		CategoryID:      3,
	}
)

//TEST GetOneQuestionById
func TestGetOneQuestionByIdSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})
	createdProblem, _ := CreateQuestion(mockDBSoal)
	oneProblem, err := GetOneQuestionById(int(createdProblem.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "1 + 9 = ?", oneProblem.Soal_pertanyaan)
		assert.Equal(t, "10", oneProblem.PilihanA)
		assert.Equal(t, "11", oneProblem.PilihanB)
		assert.Equal(t, "12", oneProblem.PilihanC)
		assert.Equal(t, "19", oneProblem.PilihanD)
		assert.Equal(t, "a", oneProblem.Solusi)
	}
}

func TestGetOneQuestionByIdError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	CreateQuestion(mockDBSoal)
	_, err := GetOneQuestionById(1)
	assert.Error(t, err)
}

//TEST EditSoal
func TestEditSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})
	createdProblem, _ := CreateQuestion(mockDBSoal)
	oneProblem, err := GetOneQuestionById(int(createdProblem.ID))
	oneProblem.Soal_pertanyaan = "100 + 19"
	oneProblem.PilihanA = "119"
	oneProblem.PilihanB = "120"
	oneProblem.PilihanC = "121"
	oneProblem.PilihanD = "122"
	oneProblem.Solusi = "a"

	editProblem, err := EditSoal(oneProblem)
	problemEdited, err := GetOneQuestionById(int(editProblem.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "100 + 19", problemEdited.Soal_pertanyaan)
		assert.Equal(t, "119", problemEdited.PilihanA)
		assert.Equal(t, "120", problemEdited.PilihanB)
		assert.Equal(t, "121", problemEdited.PilihanC)
		assert.Equal(t, "122", problemEdited.PilihanD)
		assert.Equal(t, "a", problemEdited.Solusi)
	}
}

func TestEditSoalError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	_, err := EditSoal(mockDBSoalEdit)
	assert.Error(t, err)
}

//TEST DeleteOneSoalSpecifiedId

func TestDeleteOneSoalSpecifiedIdSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})

	var message string
	var soal models.Soal
	createdProblem, _ := CreateQuestion(mockDBSoal)
	oneProblem, err := GetOneQuestionById(int(createdProblem.ID))
	_, err = DeleteOneSoalSpecifiedId(int(oneProblem.ID))

	if err := config.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&soal); err != nil {
		message = "success"
	} else {
		message = "failed"
	}

	if assert.NoError(t, err) {
		assert.Equal(t, "success", message)
	}
}

func TestDeleteOneSoalSpecifiedIdError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	_, err := DeleteOneSoalSpecifiedId(1)
	assert.Error(t, err)
}

//TEST GetAllSoalInSpecifiedCategory

func TestGetAllSoalInSpecidiedCategorySuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})

	CreateQuestion(mockDBSoal1)
	CreateQuestion(mockDBSoal2)
	CreateQuestion(mockDBSoal3)
	CreateQuestion(mockDBSoal4)

	allProblemBasedOnCategory, err := GetAllSoalInSpecifiedCategory(2)

	if assert.NoError(t, err) {
		assert.Equal(t, 2, int(allProblemBasedOnCategory[0].CategoryID))
		assert.Equal(t, 2, int(allProblemBasedOnCategory[1].CategoryID))
	}
}

func TestGetAllSoalInSpecidiedCategoryError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	_, err := GetAllSoalInSpecifiedCategory(1)
	assert.Error(t, err)
}
