package controller

import (
	"event_ticket/features/chat-AI/dto"
	"event_ticket/features/chat-AI/usecase"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AnimeRecomendationController struct {
	usecase usecase.AnimeRecomendationUsecase
}

func NewRecomendationLaptopController(usecase usecase.AnimeRecomendationUsecase) *AnimeRecomendationController {
	return &AnimeRecomendationController{usecase: usecase}
}

func (controller *AnimeRecomendationController) GetAnimeRecomendation(c echo.Context) error {
	var request dto.RequestData
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	answer, err := controller.usecase.AnimeRecomendation(request, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Adding "answer" to the response
	responseDTO := dto.Response{
		Status:         "Success",
		Recommendation: answer,
	}

	return c.JSON(http.StatusOK, responseDTO)
}
