package repository

import "github.com/google/uuid"

type User struct{
	ID uuid.UUID `json:"id"`
	Username    string `json:"username"`
	Password string `json:"password"`
	Role_id uuid.UUID `json:"role_id"`
	Created_at string `json:"created_at"`
	Ticket_id uuid.UUID `json:"ticket_id"`
}