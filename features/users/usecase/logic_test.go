package usecase

import (
	"errors"
	"event_ticket/app/mocks"
	"event_ticket/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	repoData := new(mocks.UserDataInterface)
	userUC := New(repoData)

	mockUser := users.UserCore{
		Username: "Trois",
		Email: "wawan@gmail.com", 
		Password: "123321juan",
	}

	repoData.On("Register", mockUser).Return(1, nil)
	row, err := userUC.Register(mockUser)

	assert.Nil(t, err)
	assert.Equal(t, 1, row)
}

func TestRegisterError(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	mockUser := users.UserCore{
		Email: "", 
		Password: "",
	}

	repoData.On("Register", mockUser).Return(1, nil)
	row,err := userUC.Register(mockUser)

    assert.NotNil(t,err)
    assert.Equal(t,0,row)
}

func TestRegisterFailEmailformat(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	mockUser := users.UserCore{
		Email: "juan.com", 
		Password: "juan123",
	}

	repoData.On("Register", mockUser).Return(1, nil)
	row,err := userUC.Register(mockUser)

    assert.NotNil(t,err)
    assert.Equal(t,0,row)
}

// func TestLoginSuccess(t *testing.T){
// 	repoData := new(mocks.UserDataInterface)
// 	userUC := New(repoData)

// 	expectedUser := users.UserCore{
// 		ID:       "1",
// 		Email:    "wawan@gmail.com",
// 		Username: "Trois",
// 		Password: "hashed_password",
// 	}
// 	expectedToken := "token123"

// 	repoData.On("Login", "wawan@gmail.com", "", "password").Return(expectedUser, expectedToken, nil)
// 	user, token, err := userUC.Login("wawan@gmail.com", "Trois", "password")

// 	assert.NoError(t, err)                  
// 	assert.Equal(t, expectedUser, user)     
// 	assert.Equal(t, expectedToken, token)  
	
// 	repoData.AssertCalled(t, "Login", "wawan@gmail.com", "Trois", "password")
// }

func TestLoginError(t *testing.T) {
    repoData := new(mocks.UserUseCaseInterface)
    userUC := New(repoData)

    repoData.On("Login", "wawan@gmail.com", "Trois", "password").Return(users.UserCore{}, "", errors.New("Login Failed"))

    login, token, err := userUC.Login("wawan@gmail.com", "Trois", "password")

    assert.NotNil(t, err) 
    assert.Equal(t, users.UserCore{}, login)
    assert.Equal(t, "", token)
}

func TestLoginEmptyEmailorPassword(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	user, token, err := userUC.Login("", "testuser", "password")

	assert.Error(t, err)
	assert.Equal(t, users.UserCore{}, user)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "email or password can't be empty")

	user, token, err = userUC.Login("test@example.com", "testuser", "")

	assert.Error(t, err)
	assert.Equal(t, users.UserCore{}, user)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "email or password can't be empty")
}

func TestReadAllUserSuccess(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	mockUser := []users.UserCore{
		{
			ID:       "1",
			Username: "wawan",
			Email:    "wawan@example.com",
		},
		{
			ID:       "2",
			Username: "Trois",
			Email:    "Trois@example.com",
		},
	}
	
	repoData.On("ReadAllUser").Return(mockUser, nil)
	data,err := userUC.ReadAllUser()

	assert.NoError(t, err)
	assert.Equal(t, data, mockUser)
}

func TestReadAllUserError(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	repoData.On("ReadAllUser").Return(nil, errors.New("error get data"))

	users, err := userUC.ReadAllUser()

	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Contains(t, err.Error(), "error get data")
}

func TestReadSpecificUserSuccess(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	mockUser := users.UserCore{
		ID: "abc",
		Username: "Trois",
		Email: "trois@gmail.com",
	}

	repoData.On("ReadSpecificUser","abc").Return(mockUser,nil)

	user, err := userUC.ReadSpecificUser("abc")

	assert.NoError(t, err)
	assert.Equal(t, user, mockUser)
}

func TestReadSpecificUserError(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	mockUser := users.UserCore{
		ID: "abc",
		Username: "Trois",
		Email: "trois@gmail.com",
	}

	repoData.On("ReadSpecificUser",mockUser.ID).Return(users.UserCore{},errors.New("user not found"))

	user, err := userUC.ReadSpecificUser(mockUser.ID)

	assert.Error(t, err)
	assert.Equal(t, users.UserCore{}, user)
	assert.Contains(t, err.Error(), "user not found")
}

func TestUpdateUserSuccess(t *testing.T){
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	//updated data
	userData := users.UserCore{
		ID:           userID,
		Date_of_birth: "2000-01-15",
		Phone_number:  "081234567890",
	}

	// expected data
	updatedUser := users.UserCore{
		ID:           userID,
		Date_of_birth: "2000-01-15",
		Phone_number:  "081234567890",
	}

	repoData.On("UpdateUser", userID, userData).Return(updatedUser, nil)

	user, err := userUC.UpdateUser(userID, userData)

	assert.NoError(t, err)
	assert.Equal(t, updatedUser, user)
}
func TestUpdateUserInvalidDate(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	// invalid format date
	userData := users.UserCore{
		ID:           userID,
		Date_of_birth: "15-01-2000",
		Phone_number:  "081234567890",
	}

	user, err := userUC.UpdateUser(userID, userData)

	assert.Error(t, err)
	assert.Equal(t, users.UserCore{}, user)
	assert.Contains(t, err.Error(), "Date must be in the format 'yyyy-mm-dd'")
}

func TestUpdateUserInvalidPhoneNumber(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	// invalid phone number
	userData := users.UserCore{
		ID:           userID,
		Date_of_birth: "2000-01-15",
		Phone_number:  "12345",
	}

	user, err := userUC.UpdateUser(userID, userData)

	assert.Error(t, err)
	assert.Equal(t, users.UserCore{}, user)
	assert.Contains(t, err.Error(), "Phone number format not valid")
}

func TestDeleteUserSuccess(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	repoData.On("ReadSpecificUser", userID).Return(users.UserCore{ID: userID}, nil)
	repoData.On("DeleteUser", userID).Return(nil)

	err := userUC.DeleteUser(userID)

	assert.NoError(t, err)
}

func TestDeleteUserUserNotFound(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	repoData.On("ReadSpecificUser", userID).Return(users.UserCore{}, errors.New("user not found"))

	err := userUC.DeleteUser(userID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

func TestDeleteUserError(t *testing.T) {
	repoData := new(mocks.UserUseCaseInterface)
	userUC := New(repoData)

	userID := "abc123"

	repoData.On("ReadSpecificUser", userID).Return(users.UserCore{ID: userID}, nil)
	repoData.On("DeleteUser", userID).Return(errors.New("can't delete user"))

	err := userUC.DeleteUser(userID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't delete user")
}
