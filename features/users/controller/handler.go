package controller

import (
	"event_ticket/app/middlewares"
	"event_ticket/features/users"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
		Email:         input.Email,
		Name:          input.Name,
		Username:      input.Username,
		Password:      input.Password,
		Date_of_birth: input.Date_of_birth,
		RoleId:        input.RoleId,
	}

	row, errusers := handler.userUsecase.Register(data)
	fmt.Println("row controller : ", row)
	if errusers != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errusers.Error(),
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
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   data.Email,
		"token":   token,
	})
}

func (handler *UserController) ReadAllUser(c echo.Context) error {
	userId, role := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId",
		})
	}
	if role == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get role",
		})
	}

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	data, err := handler.userUsecase.ReadAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all user",
		"data":    data,
	})
}

func (handler *UserController) ReadSpecificUser(c echo.Context) error {
	idParamstr := c.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := handler.userUsecase.ReadSpecificUser(idParams.String())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get user",
		"data":    data,
	})

}

func (handler *UserController) UpdateUser(c echo.Context) error {
	idParams := c.Param("id")

	data := new(UserRequest)
	if errBind := c.Bind(data); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	userData := users.UserCore{
		ID:            idParams,
		Name:          data.Name,
		Username:      data.Username,
		Date_of_birth: data.Date_of_birth,
		Address:       data.Address,
		Phone_number:  data.Phone_number,
	}

	fmt.Println("user data handler : ", userData)

	updatedEvent, err := handler.userUsecase.UpdateUser(idParams, userData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user updated successfully",
		"data":    updatedEvent,
	})
}

func (handler *UserController) DeleteUser(c echo.Context) error {
	idParams := c.Param("id")
	err := handler.userUsecase.DeleteUser(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}
