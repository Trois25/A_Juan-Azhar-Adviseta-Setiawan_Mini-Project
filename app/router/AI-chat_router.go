package router

import (
	"event_ticket/features/chat-AI/controller"
	"event_ticket/features/chat-AI/usecase"

	"github.com/labstack/echo/v4"
)

func SetAnimeRecomendationRoutes(e *echo.Echo) {
	AnimeRecomendationUseCase := usecase.NewUseCase()
	AnimeRecomendationController := controller.NewRecomendationLaptopController(AnimeRecomendationUseCase)

	e.POST("/anime", AnimeRecomendationController.GetAnimeRecomendation)
}
