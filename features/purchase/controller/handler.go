package controller

import (
	"event_ticket/app/middlewares"
	"event_ticket/features/purchase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type purchaseController struct {
	purchaseUsecase purchase.PurchaseUseCaseInterface
}

func New(purchaseUC purchase.PurchaseUseCaseInterface) *purchaseController {
	return &purchaseController{
		purchaseUsecase: purchaseUC,
	}
}

func (handler *purchaseController) CreatePurchase(c echo.Context) error {
	input := new(PurchaseRequest)
	errBind := c.Bind(&input)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	fmt.Println("id event ", input.EventId)
	fmt.Println("id user : ", input.UserId)
	fmt.Println("quantity event", input.Quantity)
	fmt.Println("payment status : ", input.Payment_status)

	// Validasi input
	if input.EventId == "" || input.UserId == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error, EventId and UserId can't be empty",
		})
	}

	data := purchase.PurchaseCore{
		EventId:        input.EventId,
		UserId:         input.UserId,
		Quantity:       input.Quantity,
		Payment_status: input.Payment_status,
	}

	fmt.Println("---------------------------")
	fmt.Println("id purchase : ", data.ID)
	fmt.Println("id event : ", data.EventId)
	fmt.Println("id user : ", data.UserId)
	fmt.Println("quantity event", data.Quantity)
	fmt.Println("payment status : ", data.Payment_status)
	fmt.Println("booking code : ", data.Booking_code)

	row, err := handler.purchaseUsecase.CreatePurchase(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error create purchase",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create purchase",
		"row":     row,
	})
}

func (handler *purchaseController) ReadAllPurchase(c echo.Context) error {
	data, err := handler.purchaseUsecase.ReadAllPurchase()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all purchase data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all purchase data",
		"data":    data,
	})
}

func (handler *purchaseController) ReadSpecificPurchase(c echo.Context) error {
	idParamstr := c.Param("id")

	data, err := handler.purchaseUsecase.ReadSpecificPurchase(idParamstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific purchase data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get purchase data",
		"data":    data,
	})
}

func (handler *purchaseController) UpdatePurchase(c echo.Context) error {
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

	idParams := c.Param("id")

	data := new(PurchaseRequest)
	if errBind := c.Bind(data); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	purchaseData := purchase.PurchaseCore{
		ID:             idParams,
		Payment_status: data.Payment_status,
	}

	_, err := handler.purchaseUsecase.UpdatePurchase(idParams, purchaseData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating event",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Payment Status updated",
	})
}

func (handler *purchaseController) UploadProof(c echo.Context) error {
	idParams := c.Param("id")

	data := new(PurchaseRequest)
	if errBind := c.Bind(data); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	image, err := c.FormFile("proof_image") // Pastikan ini sesuai dengan nama field file di form
	if err != nil {
		if err == http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	purchaseData := purchase.PurchaseCore{
		ID:          idParams,
		Proof_image: data.Proof_image,
	}

	_, err = handler.purchaseUsecase.UploadProof(idParams, purchaseData, image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading proof of payment",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Proof of Payment uploaded",
	})
}

func (handler *purchaseController) DeletePurchase(c echo.Context) error {
	idParams := c.Param("id")
	fmt.Println(idParams)
	err := handler.purchaseUsecase.DeletePurchase(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting purchase data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Purchase data deleted",
	})
}
