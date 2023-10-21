package router

import (
	m "event_ticket/app/middlewares"
	"event_ticket/features/roles/controller"
	"event_ticket/features/roles/database"
	"event_ticket/features/roles/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	roleRepository := database.New(db)       //menghubungkan data repo ke db
	roleUsecase := usecase.New(roleRepository) //data pada usecare berdaarkan repository
	roleController := controller.New(roleUsecase)

	e.GET("/roles", roleController.ReadAllRole, m.JWTMiddleware())
	e.POST("/roles", roleController.CreateRole)
	e.DELETE("/roles/:id", roleController.DeleteRole, m.JWTMiddleware())
}
