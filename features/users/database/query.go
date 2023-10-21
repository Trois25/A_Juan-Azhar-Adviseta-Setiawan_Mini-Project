package database

import (
	"event_ticket/app/middlewares"
	"event_ticket/features/repository"
	"event_ticket/features/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Register implements users.UserDataInterface.
func (userRep *userRepository) Register(data users.UserCore) (row int, err error) {

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0,UUIDerr
	}
	
	hashPassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	var input = repository.Users{
		ID: newUUID,
		Email:    data.Email,
		Name:     data.Name,
		Username: data.Username,
		Password: string(hashPassword),
	}

	erruser := userRep.db.Save(&input)
	if erruser.Error != nil {
		return 0, erruser.Error
	}

	return 1, nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userRepository{
		db: db,
	}
}
