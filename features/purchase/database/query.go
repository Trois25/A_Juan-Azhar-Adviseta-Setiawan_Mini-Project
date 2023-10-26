package database

import (
	"errors"
	"event_ticket/features/purchase"
	"event_ticket/features/repository"
	"event_ticket/features/storage"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type purchaseRepository struct {
	db *gorm.DB
}

// UploadProof implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) UploadProof(id string, data purchase.PurchaseCore, image *multipart.FileHeader) (purchases purchase.PurchaseCore, err error) {
	var purchaseData repository.Purchase
	errData := purchaseRepo.db.Where("id = ?", id).First(&purchaseData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return purchase.PurchaseCore{}, errors.New("purchase data not found")
		}
		return purchase.PurchaseCore{}, errData
	}

	// Pastikan UserId yang sesuai dengan referensi ke pengguna ada
	var user repository.Users
	errUser := purchaseRepo.db.Where("id = ?", purchaseData.UserId).First(&user).Error
	if errUser != nil {
		return purchase.PurchaseCore{}, errors.New("associated user not found")
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	imageURL, uploadErr := storage.UploadProofOfPayment(image)
	if uploadErr != nil {
		return purchase.PurchaseCore{}, uploadErr
	}

	purchaseData.ID = uuidID
	purchaseData.Proof_image = imageURL
	purchaseData.UpdatedAt = data.UpdatedAt

	var update = repository.Purchase{
		ID:             purchaseData.ID,
		Payment_status: purchaseData.Payment_status,
		UserId:         purchaseData.UserId,
		EventId:        purchaseData.EventId,
		Quantity:       purchaseData.Quantity,
		Total_price:    purchaseData.Total_price,
		Booking_code:   purchaseData.Booking_code,
		Proof_image:    purchaseData.Proof_image,
		CreatedAt:      purchaseData.CreatedAt,
		UpdatedAt:      purchaseData.UpdatedAt,
	}

	errSave := purchaseRepo.db.Save(&update)
	if errSave != nil {
		return purchase.PurchaseCore{}, errSave.Error
	}

	purchaseCore := purchase.PurchaseCore{
		ID:             purchaseData.ID.String(),
		UserId:         purchaseData.UserId.String(),
		EventId:        purchaseData.EventId.String(),
		Payment_status: purchaseData.Payment_status,
		Quantity:       purchaseData.Quantity,
		Total_price:    purchaseData.Total_price,
		Booking_code:   purchaseData.Booking_code.String(),
		Proof_image:    purchaseData.Proof_image,
		CreatedAt:      purchaseData.CreatedAt,
		UpdatedAt:      purchaseData.UpdatedAt,
	}

	return purchaseCore, nil
}

// ReadSpecificPurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) ReadSpecificPurchase(id string) (purchases purchase.PurchaseCore, err error) {
	var purchaseData repository.Purchase
	errData := purchaseRepo.db.Where("id = ?", id).First(&purchaseData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return purchase.PurchaseCore{}, errors.New("event not found")
		}
		return purchase.PurchaseCore{}, errData
	}

	purchaseCore := purchase.PurchaseCore{
		ID:             purchaseData.ID.String(),
		EventId:        purchaseData.EventId.String(),
		UserId:         purchaseData.UserId.String(),
		Quantity:       purchaseData.Quantity,
		Ticket_price:   purchaseData.Event.Price,
		Total_price:    purchaseData.Total_price,
		Booking_code:   purchaseData.Booking_code.String(),
		Payment_status: purchaseData.Payment_status,
		CreatedAt:      purchaseData.CreatedAt,
		UpdatedAt:      purchaseData.UpdatedAt,
	}

	return purchaseCore, nil
}

// CreatePurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) CreatePurchase(data purchase.PurchaseCore) (row int, err error) {
	fmt.Println("paling atas")
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	bookingCode, UUIDerror := uuid.NewRandom()
	if UUIDerror != nil {
		return 0, UUIDerror
	}

	var event repository.Events
	if err := purchaseRepo.db.Where("id = ?", data.EventId).First(&event).Error; err != nil {
		return 0, err
	}

	totalPrice := data.Quantity * int(event.Price)

	eventID, parseErr := uuid.Parse(data.EventId)
	if parseErr != nil {
		return 0, fmt.Errorf("error parsing EventId: %v", parseErr)
	}

	userID, parseErr := uuid.Parse(data.UserId)
	if parseErr != nil {
		return 0, fmt.Errorf("error parsing EventId: %v", parseErr)
	}
	fmt.Println(eventID, userID)

	var input = repository.Purchase{
		ID:             newUUID,
		EventId:        eventID,
		UserId:         userID,
		Quantity:       data.Quantity,
		Total_price:    float64(totalPrice),
		Booking_code:   bookingCode,
		Payment_status: data.Payment_status,
	}

	errData := purchaseRepo.db.Save(&input)
	if errData != nil {
		return 0, errData.Error
	}

	return 1, nil
}

// DeletePurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) DeletePurchase(id string) (err error) {
	var checkId repository.Purchase

	errData := purchaseRepo.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// ReadAllPurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) ReadAllPurchase() ([]purchase.PurchaseCore, error) {
	var dataPurchase []repository.Purchase

	errData := purchaseRepo.db.Find(&dataPurchase).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]purchase.PurchaseCore, len(dataPurchase))
	for i, value := range dataPurchase {
		mapData[i] = purchase.PurchaseCore{
			ID:             value.ID.String(),
			EventId:        value.EventId.String(),
			UserId:         value.UserId.String(),
			Quantity:       value.Quantity,
			Ticket_price:   value.Event.Price,
			Total_price:    value.Total_price,
			Booking_code:   value.Booking_code.String(),
			Payment_status: value.Payment_status,
			CreatedAt:      value.CreatedAt,
			UpdatedAt:      value.UpdatedAt,
		}
	}
	return mapData, nil
}

// UpdatePurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) UpdatePurchase(id string, data purchase.PurchaseCore) (purchases purchase.PurchaseCore, err error) {
	var purchaseData repository.Purchase
	errData := purchaseRepo.db.Where("id = ?", id).First(&purchaseData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return purchase.PurchaseCore{}, errors.New("purchase data not found")
		}
		return purchase.PurchaseCore{}, errData
	}

	// Pastikan UserId yang sesuai dengan referensi ke pengguna ada
	var user repository.Users
	errUser := purchaseRepo.db.Where("id = ?", purchaseData.UserId).First(&user).Error
	if errUser != nil {
		return purchase.PurchaseCore{}, errors.New("associated user not found")
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	purchaseData.ID = uuidID
	purchaseData.Payment_status = data.Payment_status
	purchaseData.UpdatedAt = data.UpdatedAt

	var update = repository.Purchase{
		ID:             purchaseData.ID,
		Payment_status: purchaseData.Payment_status,
		UserId:         purchaseData.UserId,
		EventId:        purchaseData.EventId,
		Quantity:       purchaseData.Quantity,
		Total_price:    purchaseData.Total_price,
		Booking_code:   purchaseData.Booking_code,
		CreatedAt:      purchaseData.CreatedAt,
		UpdatedAt:      purchaseData.UpdatedAt,
	}

	errSave := purchaseRepo.db.Save(&update)
	if errSave != nil {
		return purchase.PurchaseCore{}, errSave.Error
	}

	purchaseCore := purchase.PurchaseCore{
		ID:             purchaseData.ID.String(),
		UserId:         purchaseData.UserId.String(),
		EventId:        purchaseData.EventId.String(),
		Payment_status: purchaseData.Payment_status,
		Quantity:       purchaseData.Quantity,
		Total_price:    purchaseData.Total_price,
		Booking_code:   purchaseData.Booking_code.String(),
		CreatedAt:      purchaseData.CreatedAt,
		UpdatedAt:      purchaseData.UpdatedAt,
	}

	return purchaseCore, nil
}

func New(db *gorm.DB) purchase.PurchaseDataInterface {
	return &purchaseRepository{
		db: db,
	}
}
