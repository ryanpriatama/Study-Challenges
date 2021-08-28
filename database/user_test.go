package database

import (
	"SC/config"
	"SC/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockDBUser = models.User{
		Nama:      "farras",
		Email:     "farras@gmail.com",
		Password:  "123",
		TotalPoin: 0,
		Rank:      "bronze",
		Role:      "user",
	}
	mockDBUserLogin = models.User{
		Email:    "farras@gmail.com",
		Password: "123",
	}
	mockDBUserEdit = models.User{
		Nama:      "Farras Timorremboko",
		Email:     "farrastimorremboko@gmail.com",
		Password:  "12345",
		TotalPoin: 0,
		Rank:      "bronze",
		Role:      "user",
	}

	mockUser1 = models.User{
		Nama:      "umam",
		Email:     "umam@gmail.com",
		Password:  "123",
		TotalPoin: 40,
		Rank:      "silver",
		Role:      "user",
	}
	mockUser2 = models.User{
		Nama:      "frendi",
		Email:     "frendi@gmail.com",
		Password:  "456",
		TotalPoin: 30,
		Rank:      "silver",
		Role:      "user",
	}
	mockUser3 = models.User{
		Nama:      "mail",
		Email:     "mail@gmail.com",
		Password:  "789",
		TotalPoin: 35,
		Rank:      "Bronze",
		Role:      "user",
	}
	mockUser4 = models.User{
		Nama:      "farras",
		Email:     "farras@gmail.com",
		Password:  "12345",
		TotalPoin: 10,
		Rank:      "Bronze",
		Role:      "user",
	}
)

func TestCreateUserSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdUser, err := CreateUser(mockDBUser)
	if assert.NoError(t, err) {
		assert.Equal(t, "farras", createdUser.Nama)
		assert.Equal(t, "farras@gmail.com", createdUser.Email)
		assert.Equal(t, "123", createdUser.Password)
	}
}

func TestCreateUserError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := CreateUser(mockDBUser)
	assert.Error(t, err)
}

func TestCheckEmailSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	CreateUser(mockDBUser)
	same, err := CheckSameEmail("ryan@gmail.com")
	if assert.NoError(t, err) {
		assert.Equal(t, false, same)
	}
}
func TestCheckEmailSuccessSameEmail(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	CreateUser(mockDBUser)
	same, err := CheckSameEmail("farras@gmail.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, same)
	}
}
func TestCheckEmailError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	CreateUser(mockDBUser)
	_, err := CheckSameEmail("farras@gmail.com")
	assert.Error(t, err)
}

func TestLoginUserSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdUser, _ := CreateUser(mockDBUser)
	userLogin, err := LoginUsers(createdUser.Email, createdUser.Password)
	if assert.NoError(t, err) {
		assert.Equal(t, "farras", userLogin.Nama)
		assert.Equal(t, "farras@gmail.com", userLogin.Email)
		assert.Equal(t, "123", userLogin.Password)
	}
}

func TestLoginUserErrorWrongPassword(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdUser, _ := CreateUser(mockDBUser)
	_, err := LoginUsers(createdUser.Email, "456")
	assert.Error(t, err)
}

func TestGetOneUserSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdUser, _ := CreateUser(mockDBUser)
	user, err := GetOneUser(int(createdUser.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "farras", user.Nama)
		assert.Equal(t, "farras@gmail.com", user.Email)
		assert.Equal(t, "123", user.Password)
	}
}

func TestGetOneUserError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	CreateUser(mockDBUser)
	_, err := GetOneUser(1)
	assert.Error(t, err)
}

func TestEditUserSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	createdUser, _ := CreateUser(mockDBUser)
	user, err := GetOneUser(int(createdUser.ID))
	user.Nama = "Farras Timorremboko"
	user.Email = "farrastimorremboko@gmail.com"
	user.Password = "12345"
	editUser, err := EditUser(user)
	userEdited, err := GetOneUser(int(editUser.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "Farras Timorremboko", userEdited.Nama)
		assert.Equal(t, "farrastimorremboko@gmail.com", userEdited.Email)
		assert.Equal(t, "12345", userEdited.Password)
	}
}

func TestEditUserError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := EditUser(mockDBUserEdit)
	assert.Error(t, err)
}

func TestLeaderboardsSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	CreateUser(mockUser1)
	CreateUser(mockUser2)
	CreateUser(mockUser3)
	CreateUser(mockUser4)
	leaderboards, err := Leaderboards()
	if assert.NoError(t, err) {
		assert.Equal(t, "umam", leaderboards[0].Nama)
		assert.Equal(t, "mail", leaderboards[1].Nama)
		assert.Equal(t, "frendi", leaderboards[2].Nama)
		assert.Equal(t, "farras", leaderboards[3].Nama)
	}
}

func TestLeaderboardsError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.User{})
	_, err := Leaderboards()
	assert.Error(t, err)
}
