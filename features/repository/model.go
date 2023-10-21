package repository

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Username      string    `gorm:"varchar(50);not null" json:"username"`
	Password      string    `gorm:"varchar(50);not null" json:"password"`
	Role_id       uint64    `gorm:"type:varchar(50);not null;default:2" json:"role_id"`
	Name          string    `gorm:"type:varchar(50);not null" json:"name"`
	Address       string    `json:"address"`
	Email         string    `gorm:"varchar(50);not null" json:"email"`
	Date_of_birth string    `gorm:"type:date;not null" json:"date_of_birth"`
	Phone_number  string    `json:"phone_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
	Purchase_id   string    `json:"purchase_id"`
}

type Purchase struct {
	ID             string    `json:"id"`
	Ticket_id      string    `json:"ticket_id"`
	User_id        string    `json:"user_id"`
	Quantity       int       `json:"quantity"`
	Total_price    float64   `json:"total_price"`
	Booking_code   string    `json:"booking_code"`
	Payment_status string    `json:"payment_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
}

type Roles struct {
	ID        uint64    `json:"id"`
	Role_name string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type Events struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Ticket_quantity int       `json:"ticket_quantity"`
	Price           float64   `json:"price"`
	Status          string    `json:"status"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"update_at"`
}
