package controller

import (
	// "event_ticket/app/middlewares"
	"event_ticket/features/events"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	eventUsecase events.EventsUseCaseInterface
}

func New(eventUC events.EventsUseCaseInterface) *eventController {
	return &eventController{
		eventUsecase: eventUC,
	}
}

func (handler *eventController) PostEvent(c echo.Context) error {
	input := new(EventRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := events.EventsCore{
		Title:           input.Title,
		Body:            input.Body,
		Ticket_quantity: input.Ticket_quantity,
		Price:           input.Price,
		Place:           input.Place,
		Date:            input.Date,
	}

	_, errevent := handler.eventUsecase.PostEvent(data)
	if errevent != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error post event",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success post event",
	})
}

func (handler *eventController) ReadAllEvent(c echo.Context) error {
	data, err := handler.eventUsecase.ReadAllEvent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all event",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all event",
		"data":    data,
	})
}

func (handler *eventController) ReadSpecificEvent(c echo.Context) error {
	idParamstr := c.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed parse UUID",
		})
	}

	data, err := handler.eventUsecase.ReadSpecificEvent(idParams.String())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific event",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all event",
		"data":    data,
	})
}

func (handler *eventController) UpdateEvent(c echo.Context) error {
	// role, _ := middlewares.ExtractTokenUserId(c)
	// if role != "admin" {
	// 	return c.JSON(http.StatusForbidden, map[string]interface{}{
	// 		"message": "Access denied. Only admins can update events.",
	// 	})
	// }

	idParams := c.Param("id")
	
	fmt.Println("params :" ,idParams)
	data := new(EventRequest)
	if errBind := c.Bind(data); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}
	fmt.Println("data :",data)

	eventData := events.EventsCore{
		ID:              idParams,
		Title:           data.Title,
		Body:            data.Body,
		Ticket_quantity: data.Ticket_quantity,
		Price:           data.Price,
		Place:           data.Place,
		Date:            data.Date,
	}

	updatedEvent, err := handler.eventUsecase.UpdateEvent(idParams, eventData)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Error updating event",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Event updated successfully",
        "data": updatedEvent, // Anda dapat mengirim data yang telah diperbarui sebagai respons.
    })

}

func (handler *eventController) DeleteEvent(c echo.Context) error {
	// role, _ := middlewares.ExtractTokenUserId(c)
	// if role != "admin" {
	// 	return c.JSON(http.StatusForbidden, map[string]interface{}{
	// 		"message": "Access denied. Only admins can delete events.",
	// 	})
	// }

	idParams := c.Param("id")
	err := handler.eventUsecase.DeleteEvent(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting event",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Event deleted successfully",
	})
}