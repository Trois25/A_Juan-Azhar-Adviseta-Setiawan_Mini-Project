package usecase

import (
	"errors"
	"event_ticket/features/roles"
)

type roleUsecase struct {
	roleRepository roles.RoleUseCaseInterface
}

// CreateRole implements roles.RoleUseCaseInterface.
func (roleUC *roleUsecase) CreateRole(data roles.RoleCore) (err error) {
	if data.Role_name == " "{
		return errors.New("role name can't be empty")
	}

	errRole := roleUC.roleRepository.CreateRole(data)
	if errRole != nil{
		return errors.New("can't create role")
	}

	return nil
}

// DeleteRole implements roles.RoleUseCaseInterface.
func (roleUC *roleUsecase) DeleteRole(id uint64) (err error) {
	if id == 0 {
		return errors.New("role not found")
	}
	errRole := roleUC.roleRepository.DeleteRole(id)
	if errRole != nil{
		return errors.New("can't delete role")
	}

	return nil
}

// ReadAllRole implements roles.RoleUseCaseInterface.
func (roleUC *roleUsecase) ReadAllRole() ([]roles.RoleCore, error) {
	roles, err := roleUC.roleRepository.ReadAllRole()
	if err != nil{
		return nil, errors.New("error get data")
	}

	return roles, nil
}

func New(Roleuc roles.RoleDataInterface) roles.RoleUseCaseInterface {
	return &roleUsecase{
		roleRepository: Roleuc,
	}
}
