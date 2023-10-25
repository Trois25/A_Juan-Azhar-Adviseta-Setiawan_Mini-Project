package purchase

import (
	"time"
)

type PurchaseCore struct {
	ID             string    `json:"id"`
	EventId        string    `json:"event_id"`
	UserId         string    `json:"user_id"`
	Quantity       int       `json:"quantity"`
	Total_price    float64   `json:"total_price"`
	Booking_code   string    `json:"booking_code"`
	Payment_status string    `json:"payment_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
	Ticket_price   float64   `json:"ticket_price"`
}

type PurchaseDataInterface interface {
	CreatePurchase(data PurchaseCore) (row int, err error)
	ReadAllPurchase() ([]PurchaseCore, error)
	ReadSpecificPurchase(id string) (purchases PurchaseCore, err error)
	UpdatePurchase(id string, data PurchaseCore) (purchases PurchaseCore, err error)
	DeletePurchase(id string) (err error)
}

type PurchaseUseCaseInterface interface {
	CreatePurchase(data PurchaseCore) (row int, err error)
	ReadAllPurchase() ([]PurchaseCore, error)
	ReadSpecificPurchase(id string) (purchases PurchaseCore, err error)
	UpdatePurchase(id string, data PurchaseCore) (purchases PurchaseCore, err error)
	DeletePurchase(id string) (err error)
}
