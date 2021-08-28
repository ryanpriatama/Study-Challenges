package controller

import (
	"SC/auth"
	"SC/database"
	"SC/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

//post admin controller for admin signup
func AdminSignup(c echo.Context) error {
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
	admin := models.User{}
	admin.Nama = input.Nama
	admin.Email = input.Email
	admin.Password = ourEncrypt(input.Password)
	admin.Role = "admin"
	c.Bind(&admin)
	adminAdd, err := database.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapAdmin := map[string]interface{}{
		"ID":    adminAdd.ID,
		"Name":  adminAdd.Nama,
		"Email": adminAdd.Email,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new admin",
		"customer": mapAdmin,
	})
}

//get admin profile by id controller for show admin profile
func ShowAdminProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = AdminAuthorize(id, c); err != nil {
		return err
	}

	cache := Redis()
	key := fmt.Sprintf("adminId%d", id)
	adminCache, _ := cache.Get(ctx, key).Result()
	var admin models.User

	if adminCache == "" {
		admin, err = database.GetAdminid(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot find the admin",
			})
		}
		fmt.Print("Cache!!")
		json, _ := json.Marshal(admin)
		cache.Set(ctx, key, json, time.Minute)
	}
	if adminCache != "" {
		json.Unmarshal([]byte(adminCache), &admin)
	}
	mapAdmin := map[string]interface{}{
		"ID":    admin.ID,
		"Name":  admin.Nama,
		"Email": admin.Email,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success get admin",
		"admin profile": mapAdmin,
	})
}

//put admin profile by id controller for edit admin profile
func EditAdminProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = AdminAuthorize(id, c); err != nil {
		return err
	}
	admin := database.PutAdmin(id)
	c.Bind(&admin)
	adminUpdate, err1 := database.UpdateAdmin(admin)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}

	//Cache
	cache := Redis()
	key := fmt.Sprintf("adminId%d", id)
	json, _ := json.Marshal(adminUpdate)
	cache.Set(ctx, key, json, time.Minute)

	mapAdmin := map[string]interface{}{
		"ID":    adminUpdate.ID,
		"Name":  adminUpdate.Nama,
		"Email": adminUpdate.Email,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success update profile admin",
		"update customer": mapAdmin,
	})
}

//Login for admin with matching email and password
func AdminLogin(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	admin := models.User{
		Nama:     input.Email,
		Password: ourEncrypt(input.Password),
	}
	c.Bind(&admin)
	adminlogin, err := database.AdminLoginDB(admin.Email, admin.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mapAdminlogin := map[string]interface{}{
		"ID":    adminlogin.ID,
		"Name":  adminlogin.Nama,
		"Token": adminlogin.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes login",
		"admin":   mapAdminlogin,
	})
}

// AUTHORIZATION ADMIN
func AdminAuthorize(adminId int, c echo.Context) error {
	adminAuth, err := database.GetAdminid(adminId)
	loggedInAdminId, role := auth.ExtractTokenUserId(c)
	if loggedInAdminId != adminId || adminAuth.Role != role || err != nil || adminAuth.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	return nil
}

//Admin logout with update/edit the token
func AdminLogout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("adminId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetAdminid(id)
	logout.Token = ""
	admin, err := database.UpdateAdmin(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "GOODBYE!",
		"data":    admin.Nama,
	})
}
