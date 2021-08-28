package controller

import (
	"SC/auth"
	"SC/database"
	"SC/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var ctx = context.Background()

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "",
		DB:       0,
	})
	return client
}
func ourEncrypt(plain string) string {
	bytePlain := []byte(plain)
	hashed, _ := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
	return string(hashed)
}

func UserSignup(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	if input.Nama == "" || input.Email == "" || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill name, email and password correctly",
		})
	}
	if same, _ := database.CheckSameEmail(input.Email); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email already used",
		})
	}
	addUser := models.User{}
	addUser.Nama = input.Nama
	addUser.Email = input.Email
	addUser.Password = ourEncrypt(input.Password)
	addUser.TotalPoin = 0
	addUser.Rank = "bronze"
	addUser.Role = "user"
	c.Bind(&addUser)
	user, err := database.CreateUser(addUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapUser := map[string]interface{}{
		"ID":         user.ID,
		"Name":       user.Nama,
		"Email":      user.Email,
		"Total Poin": user.TotalPoin,
		"Rank":       user.Rank,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new user added",
		"data":    mapUser,
	})
}

func UserLogin(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	userData := models.User{
		Nama:     input.Email,
		Password: ourEncrypt(input.Password),
	}
	c.Bind(&userData)

	user, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username or password is not correct",
		})
	}
	mapUserLogin := map[string]interface{}{
		"ID":    user.ID,
		"Name":  user.Nama,
		"Token": user.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome",
		"users":   mapUserLogin,
	})
}

// AUTHORIZATION USER
func UserAuthorize(userId int, c echo.Context) error {
	userAuth, err := database.GetOneUser(userId)
	loggedInUserId, role := auth.ExtractTokenUserId(c)
	if loggedInUserId != userId || userAuth.Role != role || err != nil || userAuth.Role != "user" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	return nil
}
func Quote() (string, string) {
	type Response struct {
		Q string
		A string
		H string
	}
	response, _ := http.Get("https://zenquotes.io/api/random")
	responseData, _ := ioutil.ReadAll(response.Body)

	var responseObject []Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject[0].Q, responseObject[0].A
}
func ShowUserProfile(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(userId, c); err != nil {
		return err
	}
	cache := Redis()
	key := fmt.Sprintf("userId%d", userId)
	userCache, _ := cache.Get(ctx, key).Result()

	var user models.User
	if userCache == "" {
		user, err = database.GetOneUser(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot find the user",
			})
		}
		fmt.Print("Cache!!")
		json, _ := json.Marshal(user)
		cache.Set(ctx, key, json, time.Minute)

	}
	if userCache != "" {
		json.Unmarshal([]byte(userCache), &user)
	}
	quote, author := Quote()
	mapUser := map[string]interface{}{
		"ID":                  user.ID,
		"Name":                user.Nama,
		"Email":               user.Email,
		"Total Poin":          user.TotalPoin,
		"Rank":                user.Rank,
		"Random_Quote":        quote,
		"Random_Quote_Author": author,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapUser,
	})
}

func UserLogout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = UserAuthorize(id, c); err != nil {
		return err
	}
	logout, err := database.GetOneUser(id)
	logout.Token = ""
	user, err := database.EditUser(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "GOODBYE!",
		"data":    user.Nama,
	})
}

func EditUserProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(id, c); err != nil {
		return err
	}

	editUser, err := database.GetOneUser(id)
	c.Bind(&editUser)
	user, err := database.EditUser(editUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}

	//Cache
	cache := Redis()
	key := fmt.Sprintf("userId%d", id)
	json, _ := json.Marshal(user)
	cache.Set(ctx, key, json, time.Minute)

	mapUser := map[string]interface{}{
		"ID":         user.ID,
		"Name":       user.Nama,
		"Email":      user.Email,
		"Total Poin": user.TotalPoin,
		"Rank":       user.Rank,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile Updated!",
		"data":    mapUser,
	})
}

//Get top 10 player in leader board
func ShowLeaderboards(c echo.Context) error {
	users, err := database.Leaderboards()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var usersedit []map[string]interface{}
	for i := 0; i < len(users); i++ {
		mapUser := map[string]interface{}{
			"Name":       users[i].Nama,
			"Total Poin": users[i].TotalPoin,
			"Rank":       users[i].Rank,
		}
		usersedit = append(usersedit, mapUser)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get top 10 player in Leader Board",
		"Leader Board": usersedit,
	})
}
