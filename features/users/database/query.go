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
		return 0, UUIDerr
	}

	hashPassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	var input = repository.Users{
		ID:            newUUID,
		Email:         data.Email,
		Name:          data.Name,
		Date_of_birth: data.Date_of_birth,
		Username:      data.Username,
		Password:      string(hashPassword),
		RoleId:        data.RoleId,
	}

	erruser := userRep.db.Save(&input)
	if erruser.Error != nil {
		return 0, erruser.Error
	}

	return 1, nil
}

// Login implements users.UserDataInterface.
func (userRep *userRepository) Login(email string, username string, password string) (users.UserCore, string, error) {
	var data repository.Users

	tx := userRep.db.Where("email = ? OR username = ? AND password = ?", email, username, password).First(&data)
	if tx.Error != nil {
		return users.UserCore{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(data.ID, data.Email)
		if errToken != nil {
			return users.UserCore{}, "", errToken
		}
	}

	var check = users.UserCore{
		ID:       data.ID,
		Email:    data.Email,
		Password: data.Password,
		RoleId:   data.RoleId,
	}

	return check, token, nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userRepository{
		db: db,
	}
}
