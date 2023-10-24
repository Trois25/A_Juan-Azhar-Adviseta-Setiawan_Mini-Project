package usecase

import (
	"errors"
	"event_ticket/features/events"
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
func (eventUC *eventUsecase) PostEvent(data events.EventsCore) (row int, err error) {
	if data.Title == "" || data.Body == "" || data.Place == "" {
		return 0, errors.New("error, Title, Body and place can't be empty")
	}

	if data.Price < 0 || data.Ticket_quantity < 0 {
		return 0, errors.New("error, Ticket and Price must be a positive integer")
	}

	if _, parseErr := time.Parse("2006-01-02", data.Date); parseErr != nil {
		return 0, errors.New("error, Date must be in the format 'yyyy-mm-dd'")
	}

	errevents, _ := eventUC.eventRepository.PostEvent(data)
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
func (eventUC *eventUsecase) UpdateEvent(id string, data events.EventsCore) (event events.EventsCore, err error) {
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

	updatedEvent, err := eventUC.eventRepository.UpdateEvent(id, data)
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
