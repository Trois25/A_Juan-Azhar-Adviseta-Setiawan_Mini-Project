package controller

import (
	// "event_ticket/app/middlewares"
	// "event_ticket/app/middlewares"
	"event_ticket/app/middlewares"
	"event_ticket/features/roles"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type roleController struct {
	roleUsecase roles.RoleUseCaseInterface
}

func New(roleUC roles.RoleUseCaseInterface) *roleController {
	return &roleController{
		roleUsecase: roleUC,
	}
}

func (handler *roleController) CreateRole(c echo.Context) error {
	userId,role := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId",
		})
	}
	if role  == ""{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get role",
		})
	}

	if role != "admin"{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}
	
	input := new(RoleRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := roles.RoleCore{
		Role_name: input.Role_name,
	}

	errrole := handler.roleUsecase.CreateRole(data)
	if errrole != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success insert data",
	})
}

func (handler *roleController) ReadAllRole(c echo.Context) error {

	data, err := handler.roleUsecase.ReadAllRole()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all data",
		"data":    data,
	})
}
func (handler *roleController) ReadSpecificRole(c echo.Context) error {
	idParams := c.Param("id")

	data, err := handler.roleUsecase.ReadSpecificRole(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific role",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get role",
		"data":    data,
	})
}

func (handler *roleController) DeleteRole(c echo.Context) error {

	userId,role := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId",
		})
	}
	if role  == ""{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get role",
		})
	}

	if role != "admin"{
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := c.Param("id")
	if idParams == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "id can't empty",
		})
	}

	// if userId == idParams{
		number, _ := strconv.ParseUint(string(idParams), 10, 64)
		err := handler.roleUsecase.DeleteRole(number)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "failed delete role",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success delete data",
		})

}
