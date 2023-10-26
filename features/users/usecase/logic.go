package usecase

import (
	"errors"
	"event_ticket/app/middlewares"
	"event_ticket/features/users"
	"fmt"
	"time"

	"regexp"
)

type userUsecase struct {
	userRepository users.UserDataInterface
}

// Register implements users.UserUseCaseInterface.
func (uc *userUsecase) Register(data users.UserCore) (row int, err error) {
	if data.Email == "" || data.Password == "" {
		return 0, errors.New("error, email or password can't be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, data.Email)
	if !match {
		return 0, errors.New("error. email format not valid")
	}

	erruser, _ := uc.userRepository.Register(data)
	return erruser, nil
}

// Login implements users.UserUseCaseInterface.
func (uc *userUsecase) Login(email string, username string, password string) (users.UserCore, string, error) {
	if email == "" || password == "" {
		return users.UserCore{}, "", errors.New("error, email or password can't be empty")
	}

	logindata, token, err := uc.userRepository.Login(email, username, password)

	if err != nil {
		// Handle the error from the repository
		return users.UserCore{}, "", err
	}

	if middlewares.CheckPasswordHash(logindata.Password, password) {
		if err != nil {
			return users.UserCore{}, "", err
		}

		return logindata, token, nil
	}

	return users.UserCore{}, "", errors.New("Login Failed")
}


// ReadAllUser implements users.UserUseCaseInterface.
func (userUC *userUsecase) ReadAllUser() ([]users.UserCore, error) {
	users, err := userUC.userRepository.ReadAllUser()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return users, nil
}

// ReadSpecificUser implements users.UserUseCaseInterface.
func (userUC *userUsecase) ReadSpecificUser(id string) (user users.UserCore, err error) {
	if id == "" {
		return users.UserCore{}, errors.New("event ID is required")
	}

	user, err = userUC.userRepository.ReadSpecificUser(id)
	if err != nil {
		return users.UserCore{}, err
	}

	return user,nil
}

// UpdateUser implements users.UserUseCaseInterface.
func (userUC *userUsecase) UpdateUser(id string, data users.UserCore) (user users.UserCore, err error) {
	if _, parseErr := time.Parse("2006-01-02", data.Date_of_birth); parseErr != nil {
		return users.UserCore{}, errors.New("error, Date must be in the format 'yyyy-mm-dd'")
	}

	phoneRegex := `^(?:\+62|0)[0-9-]+$`
	match, _ := regexp.MatchString(phoneRegex, data.Phone_number)
	if !match {
		return users.UserCore{}, errors.New("error. Phone number format not valid")
	}
	
	fmt.Println("data logic : ",data)
	updateUser, err := userUC.userRepository.UpdateUser(id,data)
	if err != nil {
        return users.UserCore{}, err
    }

	return updateUser,nil
}

// DeleteUser implements users.UserUseCaseInterface.
func (userUC *userUsecase) DeleteUser(id string) (err error) {
	if id == "" {
		return errors.New("user data not found")
	}

	errPurchase := userUC.userRepository.DeleteUser(id)
	if errPurchase != nil {
		return errors.New("can't delete user")
	}

	return nil
}

func New(UserUcase users.UserDataInterface) users.UserUseCaseInterface {
	return &userUsecase{
		userRepository: UserUcase,
	}
}
