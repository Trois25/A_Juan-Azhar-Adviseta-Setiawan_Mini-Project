package users

import (
	"time"

	"github.com/google/uuid"
)

type UserCore struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	RoleId        uint64    `json:"role_id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Date_of_birth string    `json:"date_of_birth"`
	Phone_number  uuid.UUID `json:"phone_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
	PurchaseId    uuid.UUID `json:"purchase_id"`
	Token         string    `json:"token"`
}

type UserDataInterface interface {
	Register(data UserCore) (row int, err error)
	Login(email, username, password string) (UserCore, string, error)
	// GetData(ID uuid.UUID) (UserCore, error)
}

type UserUseCaseInterface interface {
	Register(data UserCore) (row int, err error)
	Login(email, username, password string) (UserCore, string, error)
}
