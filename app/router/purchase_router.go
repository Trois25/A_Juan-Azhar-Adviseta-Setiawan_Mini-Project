package router

import (
	m "event_ticket/app/middlewares"
	"event_ticket/features/purchase/controller"
	"event_ticket/features/purchase/database"
	"event_ticket/features/purchase/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPurchaseRouter(db *gorm.DB, e *echo.Echo) {
	purchaseRepository := database.New(db)
	purchaseUsecase := usecase.New(purchaseRepository)
	purchaseController := controller.New(purchaseUsecase)

	e.GET("/purchases", purchaseController.ReadAllPurchase, m.JWTMiddleware())
	e.GET("/purchase/:id", purchaseController.ReadSpecificPurchase, m.JWTMiddleware())
	e.POST("/purchase", purchaseController.CreatePurchase, m.JWTMiddleware())
	e.PUT("/purchase/:id", purchaseController.UpdatePurchase, m.JWTMiddleware())
	e.PUT("/purchase/proof/:id", purchaseController.UploadProof, m.JWTMiddleware())
	e.DELETE("/purchase/:id", purchaseController.DeletePurchase, m.JWTMiddleware())
}