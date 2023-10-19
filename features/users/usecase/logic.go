package usecase

import (
	"errors"
	"event_ticket/app/middlewares"
	"event_ticket/features/users"

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
		return 0, errors.New("error. format email tidak valid")
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

func New(UserUcase users.UserDataInterface) users.UserUseCaseInterface {
	return &userUsecase{
		userRepository: UserUcase,
	}
}
