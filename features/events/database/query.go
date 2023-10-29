package database

import (
	"errors"
	"event_ticket/features/events"
	"event_ticket/features/repository"
	"event_ticket/features/storage"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type eventsRepository struct {
	db *gorm.DB
}

// ReadSpecificEvent implements events.EventsDataInterface.
func (eventRepo *eventsRepository) ReadSpecificEvent(id string) (event events.EventsCore, err error) {
	var eventData repository.Events
	errData := eventRepo.db.Where("id = ?", id).First(&eventData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return events.EventsCore{}, errors.New("event not found")
		}
		return events.EventsCore{}, errData
	}

	eventCore := events.EventsCore{
		ID:              eventData.ID.String(),
		Poster_image:    eventData.Poster_image,
		Title:           eventData.Title,
		Body:            eventData.Body,
		Date:            eventData.Date,
		Price:           eventData.Price,
		Ticket_quantity: eventData.Ticket_quantity,
		Place:           eventData.Place,
		CreatedAt:       eventData.CreatedAt,
		UpdatedAt:       eventData.UpdatedAt,
	}

	return eventCore, nil
}

// PostEvent implements events.EventsDataInterface.
func (eventRepo *eventsRepository) PostEvent(data events.EventsCore, image *multipart.FileHeader) (row int, err error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	imageURL, uploadErr := storage.UploadPoster(image)
	if uploadErr != nil {
		return 0, uploadErr
	}

	if uploadErr != nil {
		return 0, uploadErr
	}

	var input = repository.Events{
		ID:              newUUID,
		Poster_image:    imageURL,
		Title:           data.Title,
		Body:            data.Body,
		Date:            data.Date,
		Price:           data.Price,
		Ticket_quantity: data.Ticket_quantity,
		Place:           data.Place,
	}

	errData := eventRepo.db.Save(&input)
	if errData != nil {
		return 0, errData.Error
	}

	return 1, nil
}

// ReadAllEvent implements events.EventsDataInterface.
func (eventRepo *eventsRepository) ReadAllEvent() ([]events.EventsCore, error) {
	var dataEvents []repository.Events

	errData := eventRepo.db.Find(&dataEvents).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]events.EventsCore, len(dataEvents))
	for i, value := range dataEvents {
		mapData[i] = events.EventsCore{
			ID:              value.ID.String(),
			Poster_image:    value.Poster_image,
			Title:           value.Title,
			Body:            value.Body,
			Date:            value.Date,
			Price:           value.Price,
			Ticket_quantity: value.Ticket_quantity,
			Place:           value.Place,
			CreatedAt:       value.CreatedAt,
			UpdatedAt:       value.UpdatedAt,
		}
	}
	return mapData, nil
}

// UpdateEvent implements events.EventsDataInterface.
func (eventRepo *eventsRepository) UpdateEvent(id string, data events.EventsCore, image *multipart.FileHeader) (event events.EventsCore, err error) {
	var eventData repository.Events
	errData := eventRepo.db.Where("id = ?", id).First(&eventData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return events.EventsCore{}, errors.New("event not found")
		}
		return events.EventsCore{}, errData
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return events.EventsCore{}, err
	}

	imageURL, uploadErr := storage.UploadPoster(image)
	if uploadErr != nil {
		return events.EventsCore{}, uploadErr
	}

	if uploadErr != nil {
		return events.EventsCore{}, uploadErr
	}

	eventData.ID = uuidID
	eventData.Poster_image = data.Poster_image
	eventData.Title = data.Title
	eventData.Body = data.Body
	eventData.Date = data.Date
	eventData.Price = data.Price
	eventData.Ticket_quantity = data.Ticket_quantity
	eventData.Place = data.Place
	eventData.UpdatedAt = data.UpdatedAt

	var update = repository.Events{
		ID:              eventData.ID,
		Poster_image:    imageURL,
		Title:           eventData.Title,
		Body:            eventData.Body,
		Date:            eventData.Date,
		Price:           eventData.Price,
		Ticket_quantity: eventData.Ticket_quantity,
		Place:           eventData.Place,
	}

	errSave := eventRepo.db.Save(&update)
	if errData != nil {
		return events.EventsCore{}, errSave.Error
	}

	eventCore := events.EventsCore{
		ID:              eventData.ID.String(),
		Poster_image:    imageURL,
		Title:           eventData.Title,
		Body:            eventData.Body,
		Date:            eventData.Date,
		Price:           eventData.Price,
		Ticket_quantity: eventData.Ticket_quantity,
		Place:           eventData.Place,
		CreatedAt:       eventData.CreatedAt,
		UpdatedAt:       eventData.UpdatedAt,
	}

	return eventCore, nil

}

// DeleteEvents implements events.EventsDataInterface.
func (eventRepo *eventsRepository) DeleteEvent(id string) (err error) {
	var checkId repository.Events

	dataId := eventRepo.db.Where("id = ?", id).First(&checkId)
	if dataId != nil {
		return errors.New("event not found")
	}

	errData := eventRepo.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func New(db *gorm.DB) events.EventsDataInterface {
	return &eventsRepository{
		db: db,
	}
}
