package usecase

import (
	"errors"
	"event_ticket/features/purchase"
)

type purchaseUsecase struct {
	purchaseRepository purchase.PurchaseUseCaseInterface
}

// CreatePurchase implements purchase.PurchaseUseCaseInterface.
func (purchaseUC *purchaseUsecase) CreatePurchase(data purchase.PurchaseCore) (row int, err error) {
	if data.UserId == "" || data.EventId == "" {
		return 0, errors.New("error, Credential data userID and eventID can't empty")
	}

	if data.Quantity <= 0 {
		return 0, errors.New("error, quantity must be a positive integer")
	}

	errPurchase, _ := purchaseUC.purchaseRepository.CreatePurchase(data)
	return errPurchase, nil
}

// DeletePurchase implements purchase.PurchaseUseCaseInterface.
func (purchaseUC *purchaseUsecase) DeletePurchase(id string) (err error) {
	if id == "" {
		return errors.New("purchase data not found")
	}

	errPurchase := purchaseUC.purchaseRepository.DeletePurchase(id)
	if errPurchase != nil {
		return errors.New("can't delete purchase")
	}

	return nil
}

// ReadAllPurchase implements purchase.PurchaseUseCaseInterface.
func (purchaseUC *purchaseUsecase) ReadAllPurchase() ([]purchase.PurchaseCore, error) {
	purchases, err := purchaseUC.purchaseRepository.ReadAllPurchase()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return purchases, nil
}

// ReadSpecificPurchase implements purchase.PurchaseUseCaseInterface.
func (purchaseUC *purchaseUsecase) ReadSpecificPurchase(id string) (purchases purchase.PurchaseCore, err error) {
	if id == "" {
		return purchase.PurchaseCore{}, errors.New("purchase ID is required")
	}

	// Call the eventRepository's ReadSpecificEvent method
	purchases, err = purchaseUC.purchaseRepository.ReadSpecificPurchase(id)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	// Check if the purchases is found in the repository, if not return an error
	if purchases.ID == "" {
		return purchase.PurchaseCore{}, errors.New("purchase data not found")
	}

	return purchases, nil
}

// UpdatePurchase implements purchase.PurchaseUseCaseInterface.
func (purchaseUC *purchaseUsecase) UpdatePurchase(id string, data purchase.PurchaseCore) (purchases purchase.PurchaseCore, err error) {
	// if id == "" {
	// 	return purchase.PurchaseCore{}, errors.New("error, Purchase ID is required")
	// }

	// if data.Payment_status == ""{
	// 	return purchase.PurchaseCore{}, errors.New("error, payment status is required")
	// }

	// updatedPurchase, err := purchaseUC.purchaseRepository.UpdatePurchase(id,data)
    // if err != nil {
    //     return purchase.PurchaseCore{}, err
    // }

	// return updatedPurchase, nil

	if id == "" {
		return purchase.PurchaseCore{}, errors.New("error, Purchase ID is required")
	}

	// Membuat objek PurchaseCore baru hanya dengan Payment_status yang diisi
	paymentStatusData := purchase.PurchaseCore{
		Payment_status: data.Payment_status,
	}

	updatedPurchase, err := purchaseUC.purchaseRepository.UpdatePurchase(id, paymentStatusData)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	return updatedPurchase, nil
}

func New(Purchaseuc purchase.PurchaseDataInterface) purchase.PurchaseUseCaseInterface {
	return &purchaseUsecase{
		purchaseRepository: Purchaseuc,
	}
}
