package repository

import "time"

type Users struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Role_id       string    `json:"role_id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Date_of_birth string    `json:"date_of_birth"`
	Phone_number  string    `json:"phone_number"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"update_at"`
	Purchase_id   string    `json:"purchase_id"`
}

type Purchase struct {
	ID             string  `json:"id"`
	Ticket_id      string  `json:"ticket_id"`
	User_id        string  `json:"user_id"`
	Quantity       int     `json:"quantity"`
	Total_price    float64 `json:"total_price"`
	Booking_code   string  `json:"booking_code"`
	Payment_status string  `json:"payment_status"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"update_at"`
}

type Roles struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	Role_name  string `json:"role_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
}

type Events struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	Body            string  `json:"body"`
	Ticket_quantity int     `json:"ticket_quantity"`
	Price           float64 `json:"price"`
	Status          string  `json:"status"`
	Date            string  `json:"date"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"update_at"`
}
