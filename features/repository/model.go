package repository

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Username      string    `gorm:"varchar(50);not null" json:"username"`
	Password      string    `gorm:"varchar(50);not null" json:"password"`
	RoleId        uint64    `gorm:"not null" json:"role_id"`
	Name          string    `gorm:"type:varchar(50);not null" json:"name"`
	Address       string    `json:"address"`
	Email         string    `gorm:"varchar(50);not null" json:"email"`
	Date_of_birth string    `gorm:"type:date;not null" json:"date_of_birth"`
	Phone_number  string    `json:"phone_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
	Role          Roles
	Purchases     []Purchase `gorm:"foreignKey:UserId"`
}

type Purchase struct {
	ID             uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	EventId        uuid.UUID `gorm:"type:varchar(50);not null" json:"event_id"`
	UserId         uuid.UUID `gorm:"type:varchar(50);not null" json:"user_id"`
	Quantity       int       `json:"quantity"`
	Total_price    float64   `json:"total_price"`
	Booking_code   uuid.UUID `json:"booking_code"`
	Proof_image    string    `json:"proof_image"`
	Payment_status string    `gorm:"default:'pending'" json:"payment_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
	Event          Events
}

type Roles struct {
	ID        uint64    `gorm:"not null;" json:"id"`
	Role_name string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type Events struct {
	ID              uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Poster_image    string    `json:poster_image`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Ticket_quantity int       `json:"ticket_quantity"`
	Price           float64   `json:"price"`
	Place           string    `json:"place"`
	Date            string    `gorm:"type:date;not null" json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"update_at"`
}
