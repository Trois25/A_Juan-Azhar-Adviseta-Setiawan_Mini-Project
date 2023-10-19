package roles

import "time"

type RoleCore struct {
	ID        uint64    `json:"id"`
	Role_name string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleDataInterface interface {
	CreateRole(data RoleCore) (err error)
	ReadAllRole() ([]RoleCore, error)
	DeleteRole(id uint64) (err error)
}

type RoleUseCaseInterface interface {
	CreateRole(data RoleCore) (err error)
	ReadAllRole() ([]RoleCore, error)
	DeleteRole(id uint64) (err error)
}
