package database

import (
	"SC/auth"
	"SC/config"
	"SC/models"
)

func CheckSameEmail(email string) (bool, error) {
	var user models.User
	if err := config.DB.Raw("select * from users where email = ?", email).Scan(&user).Error; err != nil {
		return true, err
	}
	if user.Email == email {
		return true, nil
	}
	return false, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func LoginUsers(email, password string) (models.User, error) {
	var err error
	var user models.User
	if err = config.DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = auth.CreateToken(int(user.ID))
	if err != nil {
		return user, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func EditUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Get top 10 users order by poin
func Leaderboards() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Order("total_poin desc").Limit(10).Where("role = ?", "user").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
