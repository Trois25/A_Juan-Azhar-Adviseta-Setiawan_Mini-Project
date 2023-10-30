package usecase

import (
	"errors"
	"event_ticket/app/mocks"
	"event_ticket/features/events"
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadSpecificEventSuccess(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	mockEvent := events.EventsCore{
		ID:    "abc",
		Title: "UPN Fest",
		Body:  "test upn fest",
	}

	repoData.On("ReadSpecificEvent", "abc").Return(mockEvent, nil)
	event, err := eventUC.ReadSpecificEvent("abc")

	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, event, mockEvent)
}

func TestReadSpecificEventEmptyID(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	event, err := eventUC.ReadSpecificEvent("")

	// Periksa hasil
	assert.Error(t, err)
	assert.Equal(t, events.EventsCore{}, event)
	assert.Contains(t, err.Error(), "event ID is required")
}

func TestDeleteEventSuccess(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	repoData.On("DeleteEvent", "abc").Return(nil)
	err := eventUC.DeleteEvent("abc")

	assert.NoError(t, err)
}

func TestDeleteEventEmptyID(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	err := eventUC.DeleteEvent("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "event not found")
}

func TestDeleteEventRepositoryError(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	repoData.On("DeleteEvent", "abc").Return(errors.New("repository error"))
	err := eventUC.DeleteEvent("abc")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't delete event")
}

func TestPostEventSuccess(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	//posted data
	eventData := events.EventsCore{
		Title:           "UPN Fest",
		Body:            "Fest UPN test",
		Place:           "maguwohardjo",
		Date:            time.Now().Format("2006-01-02"),
		Ticket_quantity: 100,
		Price:           50000,
	}

	repoData.On("PostEvent", eventData, (*multipart.FileHeader)(nil)).Return(1, nil)
	row, err := eventUC.PostEvent(eventData, nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, row)
}

func TestPostEventInvalidData(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	// empty place
	eventData := events.EventsCore{
		Title: "Sample Event",
		Body:  "Event Description",
	}

	row, err := eventUC.PostEvent(eventData, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Title, Body and place can't be empty")
	assert.Equal(t, 0, row)
}

func TestPostEventInvalidDate(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	// Date is not valid
	eventData := events.EventsCore{
		Title:           "UPN Fest",
		Body:            "Fest UPN test",
		Place:           "maguwohardjo",
		Date:            "Invalid Date",
		Ticket_quantity: 100,
		Price:           50000,
	}

	row, err := eventUC.PostEvent(eventData, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Date must be in the format 'yyyy-mm-dd'")
	assert.Equal(t, 0, row)
}

func TestPostEventLargeImage(t *testing.T) {
	repoData := new(mocks.EventsUseCaseInterface)
	eventUC := New(repoData)

	eventData := events.EventsCore{
		Title:           "Sample Event",
		Body:            "Event Description",
		Place:           "Event Location",
		Date:            time.Now().Format("2006-01-02"),
		Ticket_quantity: 100,
		Price:           50,
	}

	// image size error
	largeImage := &multipart.FileHeader{
		Size: 11 * 1024 * 1024,
	}

	row, err := eventUC.PostEvent(eventData, largeImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "image file size should be less than 10 MB")
	assert.Equal(t, 0, row)
}

func TestReadAllEventSuccess(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	eventData := []events.EventsCore{
		{Title: "UPN FEST 1", Body: "menfestt 1"},
		{Title: "UPN FEST 2", Body: "menfestt 2"},
		{Title: "UPN FEST 3", Body: "menfestt 3"},
	}

	mockRepo.On("ReadAllEvent").Return(eventData, nil)
	events, err := eventUC.ReadAllEvent()

	assert.NoError(t, err)
	assert.Equal(t, eventData, events)
}

func TestReadAllEventError(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	mockRepo.On("ReadAllEvent").Return([]events.EventsCore{}, errors.New("some error"))
	events, err := eventUC.ReadAllEvent()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error get data")
	assert.Empty(t, events)
}

func TestUpdateEventSuccess(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	eventData := events.EventsCore{
		Title:          "UPN Fest Update",
		Body:           "Updated Event UPN Fest",
		Place:          "Updated Event Location JIY",
		Date:           time.Now().Format("2006-01-02"),
		Ticket_quantity: 100,
		Price:          25000,
	}

	eventID := "12345abc"

	mockRepo.On("UpdateEvent", eventID, eventData, (*multipart.FileHeader)(nil)).Return(eventData, nil)
	updatedEvent, err := eventUC.UpdateEvent(eventID, eventData, nil)

	assert.NoError(t, err)
	assert.Equal(t, eventData, updatedEvent)
}

func TestUpdateEventInvalidData(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	// Invalid place
	eventData := events.EventsCore{
		Title: "Updated Event",
		Body:  "Updated Event Description",
	}

	eventID := "12345abc"

	updatedEvent, err := eventUC.UpdateEvent(eventID, eventData, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Title, Body, and Place can't be empty")
	assert.Equal(t, events.EventsCore{}, updatedEvent)
}

func TestUpdateEventInvalidDate(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	//Invalid date
	eventData := events.EventsCore{
		Title:          "UPN Fest Update",
		Body:           "Updated Event UPN Fest",
		Place:          "Updated Event Location JIY",
		Date:           "Invalid Date",
		Ticket_quantity: 100,
		Price:          25000,
	}

	eventID := "12345abc"

	updatedEvent, err := eventUC.UpdateEvent(eventID, eventData, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Date must be in the format 'yyyy-mm-dd'")
	assert.Equal(t, events.EventsCore{}, updatedEvent)
}

func TestUpdateEventLargeImage(t *testing.T) {
	mockRepo := new(mocks.EventsUseCaseInterface)
	eventUC := New(mockRepo)

	eventData := events.EventsCore{
		Title:          "Updated Event",
		Body:           "Updated Event Description",
		Place:          "Updated Event Location",
		Date:           time.Now().Format("2006-01-02"),
		Ticket_quantity: 100,
		Price:          50,
	}

	eventID := "12345abc"

	// Image size error
	largeImage := &multipart.FileHeader{
		Size: 11 * 1024 * 1024,
	}

	updatedEvent, err := eventUC.UpdateEvent(eventID, eventData, largeImage)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "image file size should be less than 10 MB")
	assert.Equal(t, events.EventsCore{}, updatedEvent)
}