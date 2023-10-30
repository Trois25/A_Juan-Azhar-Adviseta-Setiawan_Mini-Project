package usecase

import (
	"errors"
	"event_ticket/features/events"
	"mime/multipart"
	"time"
)

type eventUsecase struct {
	eventRepository events.EventsUseCaseInterface
}

// ReadSpecificEvent implements events.EventsUseCaseInterface.
func (eventUC *eventUsecase) ReadSpecificEvent(id string) (event events.EventsCore, err error) {
	if id == "" {
		return events.EventsCore{}, errors.New("event ID is required")
	}

	// Call the eventRepository's ReadSpecificEvent method
	event, err = eventUC.eventRepository.ReadSpecificEvent(id)
	if err != nil {
		return events.EventsCore{}, err
	}

	// Check if the event is found in the repository, if not return an error
	if event.ID == "" {
		return events.EventsCore{}, errors.New("event not found")
	}

	return event, nil
}

// DeleteEvents implements events.EventsUseCaseInterface.
func (eventUC *eventUsecase) DeleteEvent(id string) (err error) {
	if id == "" {
		return errors.New("event not found")
	}

	errEvent := eventUC.eventRepository.DeleteEvent(id)
	if errEvent != nil {
		return errors.New("can't delete event")
	}

	return nil
}

// PostEvent implements events.EventsUseCaseInterface.
func (eventUC *eventUsecase) PostEvent(data events.EventsCore, image *multipart.FileHeader) (row int, err error) {
	if data.Title == "" || data.Body == "" || data.Place == "" {
		return 0, errors.New("error, Title, Body and place can't be empty")
	}

	if data.Price <= 0 || data.Ticket_quantity <= 0 {
		return 0, errors.New("error, Ticket and Price must be a positive integer")
	}

	if _, parseErr := time.Parse("2006-01-02", data.Date); parseErr != nil {
		return 0, errors.New("error, Date must be in the format 'yyyy-mm-dd'")
	}

	if image != nil && image.Size > 10*1024*1024 {
        return 0, errors.New("image file size should be less than 10 MB")
    }

	errevents, errPost := eventUC.eventRepository.PostEvent(data, image)
	if errPost != nil {
		return 0, errPost
	}
	return errevents, nil
}

// ReadAllEvent implements events.EventsUseCaseInterface.
func (eventUC *eventUsecase) ReadAllEvent() ([]events.EventsCore, error) {
	events, err := eventUC.eventRepository.ReadAllEvent()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return events, nil
}

// UpdateEvent implements events.EventsUseCaseInterface.
func (eventUC *eventUsecase) UpdateEvent(id string, data events.EventsCore, image *multipart.FileHeader) (event events.EventsCore, err error) {
	if id == "" {
		return events.EventsCore{}, errors.New("error, Event ID is required")
	}

	if data.Title == "" || data.Body == "" || data.Place == "" {
		return events.EventsCore{}, errors.New("error, Title, Body, and Place can't be empty")
	}

	if data.Price < 0 || data.Ticket_quantity < 0 {
		return events.EventsCore{}, errors.New("error, Ticket and Price must be a positive integer")
	}

	if _, parseErr := time.Parse("2006-01-02", data.Date); parseErr != nil {
		return events.EventsCore{}, errors.New("error, Date must be in the format 'yyyy-mm-dd'")
	}

	if image != nil && image.Size > 10*1024*1024 {
        return events.EventsCore{}, errors.New("image file size should be less than 10 MB")
    }

	updatedEvent, err := eventUC.eventRepository.UpdateEvent(id, data, image)
    if err != nil {
        return events.EventsCore{}, err
    }

	return updatedEvent, nil
}

func New(Eventuc events.EventsDataInterface) events.EventsUseCaseInterface {
	return &eventUsecase{
		eventRepository: Eventuc,
	}
}
