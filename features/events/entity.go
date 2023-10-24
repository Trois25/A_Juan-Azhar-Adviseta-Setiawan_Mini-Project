package events

import "time"

type EventsCore struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Ticket_quantity int       `json:"ticket_quantity"`
	Price           float64   `json:"price"`
	Place           string    `json:"place"`
	Date            string    `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"update_at"`
}

type EventsDataInterface interface {
	PostEvent(data EventsCore) (row int, err error)
	ReadAllEvent() ([]EventsCore, error)
	ReadSpecificEvent(id string) (event EventsCore, err error)
	UpdateEvent(id string, data EventsCore) (event EventsCore, err error)
	DeleteEvent(id string) (err error)
}

type EventsUseCaseInterface interface {
	PostEvent(data EventsCore) (row int, err error)
	ReadAllEvent() ([]EventsCore, error)
	ReadSpecificEvent(id string) (event EventsCore, err error)
	UpdateEvent(id string, data EventsCore) (event EventsCore, err error)
	DeleteEvent(id string) (err error)
}
