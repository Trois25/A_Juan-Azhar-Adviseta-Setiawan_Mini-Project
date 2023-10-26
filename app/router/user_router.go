package router

import (
	// m "event_ticket/app/middlewares"
	"event_ticket/features/users/controller"
	"event_ticket/features/users/database"
	"event_ticket/features/users/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitUserRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := database.New(db)         //menghubungkan data repo ke db
	userUsecase := usecase.New(userRepository) //data pada usecare berdaarkan repository
	userController := controller.New(userUsecase)

	e.GET("/users", userController.ReadAllUser)
	e.GET("/user/:id", userController.ReadSpecificUser)
	e.POST("/user", userController.Register)
	e.POST("/user/login", userController.Login)
	e.PUT("/user/:id", userController.UpdateUser)
	e.DELETE("/user/:id", userController.DeleteUser)

}
