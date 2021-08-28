package database

import (
	"SC/config"
	"SC/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBAdmin = models.User{
		Nama:     "bulba",
		Email:    "bulba@gmail.com",
		Password: "1111",
		Role:     "admin",
	}
	mockDBAdminLogin = models.User{
		Email:    "bulba@gmail.com",
		Password: "1111",
	}
	mockDBAdminEdit = models.User{
		Nama:     "bulbasaur",
		Email:    "bulbasaur@gmail.com",
		Password: "2222",
		Role:     "admin",
	}
)

func TestCreateAdminSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, err := CreateAdmin(mockDBAdmin)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulba", createdAdmin.Nama)
		assert.Equal(t, "bulba@gmail.com", createdAdmin.Email)
		assert.Equal(t, "1111", createdAdmin.Password)
	}
}

func TestCreateAdminError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := CreateAdmin(mockDBAdmin)
	assert.Error(t, err)
}

func TestLoginAdminSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, _ := CreateAdmin(mockDBAdmin)
	adminLogin, err := AdminLoginDB(createdAdmin.Email, createdAdmin.Password)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulba", adminLogin.Nama)
		assert.Equal(t, "bulba@gmail.com", adminLogin.Email)
		assert.Equal(t, "1111", adminLogin.Password)
	}
}

func TestLoginAdminErrorWrongPassword(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, _ := CreateAdmin(mockDBAdmin)
	_, err := AdminLoginDB(createdAdmin.Email, "456")
	assert.Error(t, err)
}

func TestGetOneAdminSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, _ := CreateAdmin(mockDBAdmin)
	admin, err := GetAdminid(int(createdAdmin.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "bulba", admin.Nama)
		assert.Equal(t, "bulba@gmail.com", admin.Email)
		assert.Equal(t, "1111", admin.Password)
	}
}

func TestGetOneAdminError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	CreateAdmin(mockDBAdmin)
	_, err := GetAdminid(1)
	assert.Error(t, err)
}

func TestEditAdminSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, _ := CreateAdmin(mockDBAdmin)
	admin, err := GetAdminid(int(createdAdmin.ID))
	admin.Nama = "bulbasaur"
	admin.Email = "bulbasaur@gmail.com"
	admin.Password = "2222"
	editAdmin, err := UpdateAdmin(admin)
	adminEdited, err := GetAdminid(int(editAdmin.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", adminEdited.Nama)
		assert.Equal(t, "bulbasaur@gmail.com", adminEdited.Email)
		assert.Equal(t, "2222", adminEdited.Password)
	}
}

func TestEditAdminError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := UpdateAdmin(mockDBAdminEdit)
	assert.Error(t, err)
}

func TestPutOneAdminIdSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdAdmin, _ := CreateAdmin(mockDBAdmin)
	admin := PutAdmin(int(createdAdmin.ID))
	assert.Equal(t, "bulba", admin.Nama)
	assert.Equal(t, "bulba@gmail.com", admin.Email)
	assert.Equal(t, "1111", admin.Password)
}
