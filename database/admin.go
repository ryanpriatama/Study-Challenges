package database

import (
	"SC/auth"
	"SC/config"
	"SC/models"
)

//create admin in tabel user in DB
func CreateAdmin(admin models.User) (models.User, error) {
	if err := config.DB.Save(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

//Get admin by id in tabel user
func GetAdminid(id int) (models.User, error) {
	var admin models.User
	var count int64
	if err1 := config.DB.Model(&admin).Where("id=?", id).Count(&count).Error; count == 0 {
		return admin, err1
	}
	if err := config.DB.Find(&admin, "id=?", id).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

//part 1 edit/put, Edit column admin by id in tabel user
func UpdateAdmin(admin models.User) (models.User, error) {
	if tx := config.DB.Save(&admin).Error; tx != nil {
		return admin, tx
	}
	return admin, nil
}

//part 2 edit/put, get 1 specified admin with models.User struct output
func PutAdmin(id int) models.User {
	var admin models.User
	config.DB.Find(&admin, "id=?", id)
	return admin
}

//Login for admin with matching email and password
func AdminLoginDB(email, password string) (models.User, error) {
	var admin models.User
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&admin).Error; err != nil {
		return admin, err
	}
	admin.Token, err = auth.CreateAdminToken(int(admin.ID))
	if err != nil {
		return admin, err
	}
	if err := config.DB.Save(admin).Error; err != nil {
		return admin, err
	}
	return admin, err
}
