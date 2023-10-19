package repository

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Username      string    `gorm:"varchar(50);not null" json:"username"`
	Password      string    `gorm:"varchar(50);not null" json:"password"`
	Role_id       uuid.UUID `gorm:"type:varchar(50);not null" json:"role_id"`
	Name          string    `gorm:"type:varchar(50);not null" json:"name"`
	Address       string    `json:"address"`
	Email         string    `gorm:"varchar(50);not null" json:"email"`
	Date_of_birth string    `json:"date_of_birth"`
	Phone_number  uuid.UUID `json:"phone_number"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"update_at"`
	Purchase_id   uuid.UUID `json:"purchase_id"`
}

type Purchase struct {
	ID             uuid.UUID `json:"id"`
	Ticket_id      uuid.UUID `json:"ticket_id"`
	User_id        uuid.UUID `json:"user_id"`
	Quantity       int       `json:"quantity"`
	Total_price    float64   `json:"total_price"`
	Booking_code   string    `json:"booking_code"`
	Payment_status string    `json:"payment_status"`
	Created_at     time.Time `json:"created_at"`
}

type Roles struct {
	ID         uuid.UUID `json:"id"`
	Role_name  string    `json:"role_name"`
	Created_at time.Time `json:"created_at"`
}

type Events struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Ticket_quantity int       `json:"ticket_quantity"`
	Price           float64   `json:"price"`
	Status          string    `json:"status"`
	Date            time.Time `json:"date"`
	Created_at      time.Time `json:"created_at"`
}
