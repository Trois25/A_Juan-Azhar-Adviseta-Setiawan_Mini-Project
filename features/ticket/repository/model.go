package repository

import "github.com/google/uuid"

type Ticket struct{
	ID uuid.UUID `json:"id"`
	Quantity    int `json:"quantity"`
	Events_id uuid.UUID `json:"Events_id"`
}