package repository

import "github.com/google/uuid"

type Profile struct{
	ID uuid.UUID `json:"id"`
	User_id uuid.UUID `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email string `json:"email"`
	Date_of_birth string `json:"date_of_birth"`
	Phone_number uuid.UUID `json:"phone_number"`
	Updated_at string `json:"update_at"`
}