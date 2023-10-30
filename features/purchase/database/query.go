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

	//find firs purchase with input id
	var purchaseData repository.Purchase
	errData := purchaseRepo.db.Where("id = ?", id).First(&purchaseData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return purchase.PurchaseCore{}, errors.New("purchase data not found")
		}
		return purchase.PurchaseCore{}, errData
	}

	// check validation userId
	var user repository.Users
	errUser := purchaseRepo.db.Where("id = ?", purchaseData.UserId).First(&user).Error
	if errUser != nil {
		return purchase.PurchaseCore{}, errors.New("associated user not found")
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	//upload image
	imageURL, uploadErr := storage.UploadProofOfPayment(image)
	if uploadErr != nil {
		return purchase.PurchaseCore{}, uploadErr
	}

	if uploadErr != nil {
		return purchase.PurchaseCore{}, uploadErr
	}

	//update data
	purchaseData.ID = uuidID
	purchaseData.Proof_image = imageURL
	purchaseData.UpdatedAt = data.UpdatedAt

	//mapping for input to purchase db
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

	//save to purchase db
	errSave := purchaseRepo.db.Save(&update)
	if errSave != nil {
		return purchase.PurchaseCore{}, errSave.Error
	}

	//mapping to core for data output
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

	//generate UUID
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	bookingCode, UUIDerror := uuid.NewRandom()
	if UUIDerror != nil {
		return 0, UUIDerror
	}
	
	//find first user with input ID
	var user repository.Users
	if err := purchaseRepo.db.Where("id = ?", data.UserId).First(&user).Error; err != nil {
		return 0, err
	}

	//find first events with input ID
	var event repository.Events
	if err := purchaseRepo.db.Where("id = ?", data.EventId).First(&event).Error; err != nil {
		return 0, err
	}
	fmt.Println("eventID : ",data.EventId)

	//check ticket and total price
	totalTicket := event.Ticket_quantity - data.Quantity
	if totalTicket <= 0 {
		return 0, errors.New("ticket is out of stock")
	}
	totalPrice := data.Quantity * int(event.Price)

	if data.UserId != user.ID.String(){
		return 0, errors.New("user is not found")
	}

	if data.EventId != event.ID.String(){
		return 0, errors.New("event is not found")
	}

	// update event
	errUpdateEvent := purchaseRepo.db.Model(&event).Update("ticket_quantity", totalTicket).Error
	if errUpdateEvent != nil {
		return 0, errUpdateEvent
	}

	//parse ID from string back to UUID
	eventID, parseErr := uuid.Parse(data.EventId)
	if parseErr != nil {
		return 0, fmt.Errorf("error parsing EventId: %v", parseErr)
	}

	userID, parseErr := uuid.Parse(data.UserId)
	if parseErr != nil {
		return 0, fmt.Errorf("error parsing EventId: %v", parseErr)
	}
	fmt.Println(eventID, userID)

	//mapping for input to db
	var input = repository.Purchase{
		ID:             newUUID,
		EventId:        eventID,
		UserId:         userID,
		Quantity:       data.Quantity,
		Total_price:    float64(totalPrice),
		Booking_code:   bookingCode,
		Payment_status: data.Payment_status,
	}

	//save purchase data
	errData := purchaseRepo.db.Save(&input)
	if errData != nil {
		return 0, errData.Error
	}

	return 1, nil
}

// DeletePurchase implements purchase.PurchaseDataInterface.
func (purchaseRepo *purchaseRepository) DeletePurchase(id string) (err error) {
	var checkId repository.Purchase

	dataId := purchaseRepo.db.Where("id = ?", id).First(&checkId)
	if dataId != nil {
		return errors.New("purchase data not found")
	}

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

	//find specific purchase data
	errData := purchaseRepo.db.Where("id = ?", id).First(&purchaseData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return purchase.PurchaseCore{}, errors.New("purchase data not found")
		}
		return purchase.PurchaseCore{}, errData
	}

	// check userId
	var user repository.Users
	errUser := purchaseRepo.db.Where("id = ?", purchaseData.UserId).First(&user).Error
	if errUser != nil {
		return purchase.PurchaseCore{}, errors.New("associated user not found")
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return purchase.PurchaseCore{}, err
	}

	//Edit data
	purchaseData.ID = uuidID
	purchaseData.Payment_status = data.Payment_status
	purchaseData.UpdatedAt = data.UpdatedAt

	//check image
	if data.Proof_image != "" {
		purchaseData.Proof_image = data.Proof_image
	}

	//map for insert to db
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

	//map data to output
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

func New(db *gorm.DB) purchase.PurchaseDataInterface {
	return &purchaseRepository{
		db: db,
	}
}
