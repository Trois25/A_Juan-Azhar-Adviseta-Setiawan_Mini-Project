package database

import (
	"errors"
	"event_ticket/features/repository"
	"event_ticket/features/roles"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

// ReadSpecificUser implements roles.RoleDataInterface.
func (roleRepo *roleRepository) ReadSpecificRole(id string) (role roles.RoleCore, err error) {
	var roleData repository.Roles
	errData := roleRepo.db.Where("id = ?", id).First(&roleData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return roles.RoleCore{}, errors.New("user not found")
		}
		return roles.RoleCore{}, errData
	}

	roleCore := roles.RoleCore{
		ID:            roleData.ID,
		Role_name: roleData.Role_name,
	}

	return roleCore, nil
}

// CreateRole implements roles.RoleDataInterface.
func (roleRepo *roleRepository) CreateRole(data roles.RoleCore) (err error) {
	var input = repository.Roles{
		Role_name: data.Role_name,
	}

	errData := roleRepo.db.Save(&input)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// DeleteRole implements roles.RoleDataInterface.
func (roleRepo *roleRepository) DeleteRole(id uint64) (err error) {
	var checkId repository.Roles

	errData := roleRepo.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// ReadAllRole implements roles.RoleDataInterface.
func (roleRepo *roleRepository) ReadAllRole() ([]roles.RoleCore, error) {
	var dataRoles []repository.Roles

	errData := roleRepo.db.Find(&dataRoles).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]roles.RoleCore, len(dataRoles))
	for i, value := range dataRoles {
		mapData[i] = roles.RoleCore{
			ID:        value.ID,
			Role_name: value.Role_name,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
	}

	return mapData, nil

}

func New(db *gorm.DB) roles.RoleDataInterface {
	return &roleRepository{
		db: db,
	}
}
