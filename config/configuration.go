package config

import (
	"SC/models"
	"os"
	"strconv"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var HTTP_PORT int

func Init_DB() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Set_soal{})
	DB.AutoMigrate(&models.Set_soal_detail{})
	DB.AutoMigrate(&models.Soal{})
	DB.AutoMigrate(&models.Category{})
}

func Init_DB_Test() {
	// connectionString := os.Getenv("CONNECTION_STRING_TEST")
	var err error
	DB, err = gorm.Open(mysql.Open("root:12345@tcp(172.17.0.1:3308)/study_test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
