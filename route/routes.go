package route

import (
	"SC/constant"
	"SC/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	// SIGNUP FOR USER AND ADMIN
	e.POST("/users/signup", controller.UserSignup)
	e.POST("/admin/signup", controller.AdminSignup)

	//LOGIN FOR USER AND ADMIN
	e.POST("/users/login", controller.UserLogin)
	e.POST("/admin/login", controller.AdminLogin)

	//AUTHORIZATION JWT
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	//LOGOUT FOR USER AND ADMIN
	eJwt.PUT("/users/:userId/logout", controller.UserLogout)
	eJwt.PUT("/admin/:adminId/logout", controller.AdminLogout)

	//USER PROFILE
	eJwt.GET("/users/:id", controller.ShowUserProfile)
	eJwt.PUT("/users/:id", controller.EditUserProfile)
	eJwt.GET("/users", controller.ShowLeaderboards)

	//ADMIN PROFILE
	eJwt.GET("/admin/:id", controller.ShowAdminProfile)
	eJwt.PUT("/admin/:id", controller.EditAdminProfile)

	//ADMIN FEATURES INPUT QUESTION
	eJwt.POST("/admin/soal", controller.SubmitQuestionAdmin)
	eJwt.PUT("/admin/soal/:soalId", controller.EditQuestion)
	eJwt.DELETE("/admin/soal/:soalId", controller.DeleteQuestion)
	eJwt.GET("/admin/soal/:soalId", controller.GetQuestionById)
	eJwt.GET("/admin/soal/mapel/:MataPelajaranId", controller.GetQuestionByCategory)

	//ADMIN FEATURE REVIEW SUBMITTED QUESTION FROM USER
	eJwt.GET("/admin/submit_soal/:kategori_materi_pelajaran_id", controller.ShowSubmittedQuestion) //SHOW QUESTIONS ARE NOTE REVIEWED BY CATEGORY
	eJwt.PUT("/admin/submit_soal/approval/:id", controller.EditSubmitQuestion)                     //APPROVAL THE QUESTION (APPROVED OR REJECT)

	//USER FEATURE SUBMIT NEW QUESTION
	eJwt.POST("/users/submit_soal", controller.SubmitQuestion)

	//USER FEATURE EXERCISE
	eJwt.POST("/users/:user_id/soal", controller.GenerateRandomQuestion)
	eJwt.GET("/users/:user_id/soal/:set_soal_id", controller.ShowActiveQuestion)
	eJwt.PUT("/users/:user_id/soal/:set_soal_id", controller.AnswerQuestion)
	eJwt.GET("/users/:user_id/soal/:set_soal_id/solution", controller.ShowSolution)
}
