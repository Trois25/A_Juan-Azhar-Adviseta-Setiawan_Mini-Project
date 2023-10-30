package router

import (
	m "event_ticket/app/middlewares"
	"event_ticket/features/events/controller"
	"event_ticket/features/events/database"
	"event_ticket/features/events/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitEventRouter(db *gorm.DB, e *echo.Echo) {
	eventRepository := database.New(db)
	eventUsecase := usecase.New(eventRepository)
	eventController := controller.New(eventUsecase)

	e.GET("/events", eventController.ReadAllEvent)
	e.GET("/event/:id", eventController.ReadSpecificEvent)
	e.POST("/event", eventController.PostEvent, m.JWTMiddleware())
	e.PUT("/event/:id", eventController.UpdateEvent, m.JWTMiddleware())
	e.DELETE("/event/:id", eventController.DeleteEvent, m.JWTMiddleware())
}
