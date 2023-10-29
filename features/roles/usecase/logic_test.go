package usecase

import (
	"errors"
	"event_ticket/app/mocks"
	"event_ticket/features/roles"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSpecificRoleSuccess(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	mockRole := roles.RoleCore{
		ID: 1,
		Role_name: "admin",
	}

	repoData.On("ReadSpecificRole","1").Return(mockRole,nil)

	role, err := roleUC.ReadSpecificRole("1")

	assert.NoError(t, err)
	assert.Equal(t, role, mockRole)
}

func TestReadSpecificRoleError(t *testing.T){
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	repoData.On("ReadSpecificRole","1").Return(roles.RoleCore{},errors.New("role not found"))

	role, err := roleUC.ReadSpecificRole("1")

	assert.Error(t, err)
	assert.Equal(t, roles.RoleCore{}, role)
	assert.Contains(t, err.Error(), "role not found")
}

func TestCreateRoleSuccess(t *testing.T){
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	mockUser := roles.RoleCore{
		ID: 1,
		Role_name: "ad,om",
	}

	repoData.On("CreateRole", mockUser).Return(nil)
	err := roleUC.CreateRole(mockUser)

	assert.Nil(t, err)
}

func TestCreateRoleError(t *testing.T){
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	roleData := roles.RoleCore{
		Role_name: "Admin",
	}

	repoData.On("CreateRole", roleData).Return(errors.New("can't create role"))

	err := roleUC.CreateRole(roleData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't create role")
}

func TestCreateRoleEmptyRoleName(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	roleData := roles.RoleCore{
		Role_name: "",
	}

	err := roleUC.CreateRole(roleData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "role name can't be empty")
}

func TestDeleteRoleSuccess(t *testing.T){
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	roleID := uint64(1)
	roleIDString := strconv.FormatUint(roleID, 10)

	repoData.On("ReadSpecificRole", roleIDString).Return(roles.RoleCore{ID: roleID}, nil)
	repoData.On("DeleteRole", roleID).Return(nil)

	err := roleUC.DeleteRole(roleID)

	assert.NoError(t, err)
}

func TestDeleteRoleNotFound(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	roleID := uint64(1)

	roleIDString := strconv.FormatUint(roleID, 10)

	repoData.On("ReadSpecificRole", roleIDString).Return(roles.RoleCore{}, errors.New("role not found"))

	err := roleUC.DeleteRole(roleID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "role not found")
}

func TestDeleteRoleError(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	roleID := uint64(1)
	roleIDString := strconv.FormatUint(roleID, 10)

	repoData.On("ReadSpecificRole", roleIDString).Return(roles.RoleCore{ID: roleID}, nil)
	repoData.On("DeleteRole", roleID).Return(errors.New("can't delete role"))

	err := roleUC.DeleteRole(roleID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't delete role")
}

func TestReadAllRoleSuccess(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	expectedRoles := []roles.RoleCore{
		{ID: 1, Role_name: "Admin"},
		{ID: 2, Role_name: "User"},
	}

	repoData.On("ReadAllRole").Return(expectedRoles, nil)
	roles, err := roleUC.ReadAllRole()

	assert.NoError(t, err)
	assert.NotNil(t, roles)
	assert.Equal(t, expectedRoles, roles)
}

func TestReadAllRoleError(t *testing.T) {
	repoData := new(mocks.RoleUseCaseInterface)
	roleUC := New(repoData)

	repoData.On("ReadAllRole").Return(nil, errors.New("error get data"))
	roles, err := roleUC.ReadAllRole()

	assert.Error(t, err)
	assert.Nil(t, roles)
	assert.Contains(t, err.Error(), "error get data")
}