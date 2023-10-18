package repository

import "github.com/google/uuid"

type Roles struct{
	ID uuid.UUID `json:"id"`
	Role_name    string `json:"role_name"`
	Created_at string `json:"created_at"`
}