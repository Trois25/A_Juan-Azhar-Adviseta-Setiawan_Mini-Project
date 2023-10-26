package database

import (
	"errors"
	"event_ticket/app/middlewares"
	"event_ticket/features/repository"
	"event_ticket/features/users"
	"fmt"

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
	var roleName string

	if tx.RowsAffected > 0 {
		var roleData repository.Roles
		if err := userRep.db.Where("ID = ?", data.RoleId).First(&roleData).Error; err != nil {
			return users.UserCore{}, "", err
		}
		roleName = roleData.Role_name

		var errToken error
		token, errToken = middlewares.CreateToken(data.ID, roleName)
		if errToken != nil {
			return users.UserCore{}, "", errToken
		}
	}

	var check = users.UserCore{
		ID:       data.ID.String(),
		Email:    data.Email,
		Password: data.Password,
		RoleId:   data.RoleId,
	}

	return check, token, nil
}

// ReadAllUser implements users.UserDataInterface.
func (userRepo *userRepository) ReadAllUser() ([]users.UserCore, error) {
	var dataUser []repository.Users

	errData := userRepo.db.Find(&dataUser).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]users.UserCore, len(dataUser))
	for i, value := range dataUser {
		mapData[i] = users.UserCore{
			ID:            value.ID.String(),
			Name:          value.Name,
			Username:      value.Username,
			Email:         value.Email,
			Date_of_birth: value.Date_of_birth,
			Address:       value.Address,
			Phone_number:  value.Phone_number,
			RoleId:        value.RoleId,
		}
	}
	return mapData, nil
}

// ReadSpecificUser implements users.UserDataInterface.
func (userRepo *userRepository) ReadSpecificUser(id string) (user users.UserCore, err error) {
	var userData repository.Users
	errData := userRepo.db.Where("id = ?", id).First(&userData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return users.UserCore{}, errors.New("user not found")
		}
		return users.UserCore{}, errData
	}

	userCore := users.UserCore{
		ID:            userData.ID.String(),
		Name:          userData.Name,
		Username:      userData.Username,
		Email:         userData.Email,
		Date_of_birth: userData.Date_of_birth,
		Address:       userData.Address,
		Phone_number:  userData.Phone_number,
		RoleId:        userData.RoleId,
	}

	return userCore, nil
}

// UpdateUser implements users.UserDataInterface.
func (userRepo *userRepository) UpdateUser(id string, data users.UserCore) (user users.UserCore, err error) {
	var userData repository.Users
	errData := userRepo.db.Where("id = ?", id).First(&userData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return users.UserCore{}, errors.New("event not found")
		}
		return users.UserCore{}, errData
	}

	var role repository.Roles
	errUser := userRepo.db.Where("id = ?", userData.RoleId).First(&role).Error
	if errUser != nil {
		return users.UserCore{}, errors.New("associated role not found")
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return users.UserCore{}, err
	}

	userData.ID = uuidID
	userData.Name = data.Name
	userData.Username = data.Username
	userData.Date_of_birth = data.Date_of_birth
	userData.Address = data.Address
	userData.Phone_number = data.Phone_number
	userData.UpdatedAt = data.UpdatedAt

	var update = repository.Users{
		ID:            userData.ID,
		Name:          userData.Name,
		Username:      userData.Username,
		Email:         userData.Email,
		Password:      userData.Password,
		Date_of_birth: userData.Date_of_birth,
		RoleId:        userData.RoleId,
		Address:       userData.Address,
		Phone_number:  userData.Phone_number,
		CreatedAt:     userData.CreatedAt,
		UpdatedAt:     userData.UpdatedAt,
	}

	fmt.Println("data query : ", update)
	errSave := userRepo.db.Save(&update)
	if errData != nil {
		return users.UserCore{}, errSave.Error
	}

	eventCore := users.UserCore{
		ID:            userData.ID.String(),
		Name:          userData.Name,
		Username:      userData.Username,
		Date_of_birth: userData.Date_of_birth,
		Address:       userData.Address,
		Phone_number:  userData.Phone_number,
		UpdatedAt:     userData.UpdatedAt,
	}

	return eventCore, nil
}

// DeleteUser implements users.UserDataInterface.
func (userRepo *userRepository) DeleteUser(id string) (err error) {
	var checkId repository.Users

	errData := userRepo.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userRepository{
		db: db,
	}
}
