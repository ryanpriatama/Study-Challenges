package database

import (
	"SC/config"
	"SC/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBSoalInput = models.Soal{
		Soal_pertanyaan: "5 x 2 = ?",
		PilihanA:        "10",
		PilihanB:        "123",
		PilihanC:        "25",
		PilihanD:        "15",
		Jawaban:         "a",
		KesulitanID:     1,
		Solusi:          "apabila 5 dikali 2, maka hasil pasti 10",
		Approval:        "accept",
		CategoryID:      4,
	}
)

func TestSubmitSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})
	soalCreated, err := CreateQuestion(mockDBSoalInput)
	if assert.NoError(t, err) {
		assert.Equal(t, "5 x 2 = ?", soalCreated.Soal_pertanyaan)
		assert.Equal(t, "10", soalCreated.PilihanA)
		assert.Equal(t, "123", soalCreated.PilihanB)
		assert.Equal(t, "25", soalCreated.PilihanC)
		assert.Equal(t, "15", soalCreated.PilihanD)
	}
}

func TestSubmitSoalError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	_, err := CreateQuestion(mockDBSoalInput)
	assert.Error(t, err)
}
