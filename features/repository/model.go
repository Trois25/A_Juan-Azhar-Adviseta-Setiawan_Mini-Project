package repository

import "github.com/google/uuid"

type Users struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role_id    uuid.UUID `json:"role_id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Date_of_birth string    `json:"date_of_birth"`
	Phone_number  uuid.UUID `json:"phone_number"`
	Created_at    string    `json:"created_at"`
	Updated_at    string    `json:"update_at"`
	Purchase_id       uuid.UUID `json:"purchase_id"`
}

type Purchase struct{
	ID        uuid.UUID `json:"id"`
	Ticket_id uuid.UUID `json:"ticket_id"`
	User_id uuid.UUID `json:"user_id"`
	Quantity  int       `json:"quantity"`
	Total_price float64 `json:"total_price"`
}

type Ticket struct {
	ID        uuid.UUID `json:"id"`
	Quantity  int       `json:"quantity"`
	Price float64 `json:"price"`
	Events_id uuid.UUID `json:"Events_id"`
}

type Roles struct {
	ID         uuid.UUID `json:"id"`
	Role_name  string    `json:"role_name"`
	Created_at string    `json:"created_at"`
}

type Events struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Status     string    `json:"status"`
	Created_at string    `json:"created_at"`
}
