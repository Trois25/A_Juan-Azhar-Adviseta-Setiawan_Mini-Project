package router

import (
	// m "event_ticket/app/middlewares"
	"event_ticket/features/users/controller"
	"event_ticket/features/users/database"
	"event_ticket/features/users/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := database.New(db)         //menghubungkan data repo ke db
	userUsecase := usecase.New(userRepository) //data pada usecare berdaarkan repository
	userController := controller.New(userUsecase)

	e.POST("/users", userController.Register)
	e.POST("/users/login", userController.Login)
	e.GET("/roles", roleController.ReadAllRole, m.JWTMiddleware())
	e.POST("/roles", roleController.CreateRole)
	e.DELETE("/roles/:id", roleController.DeleteRole, m.JWTMiddleware())
}
