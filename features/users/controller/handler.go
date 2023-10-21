package controller

import (
	"event_ticket/features/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.UserUseCaseInterface
}

func New(userUC users.UserUseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) Register(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := users.UserCore{
		Email:    input.Email,
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Date_of_birth: input.Date_of_birth,
		Role_id:  input.Role_id,
	}

	row, errusers := handler.userUsecase.Register(data)
	if errusers != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success insert data",
		"data":    row,
	})
}

func (handler *UserController) Login(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := users.UserCore{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	data, token, err := handler.userUsecase.Login(data.Email, data.Username, data.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error login",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   data.Email,
		"token":   token,
	})
}
