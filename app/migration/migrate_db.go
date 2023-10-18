package migration

import (
	"event_ticket/features/repository"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&repository.Users{})
	db.AutoMigrate(&repository.Ticket{})
	db.AutoMigrate(&repository.Roles{})
	db.AutoMigrate(&repository.Purchase{})
	db.AutoMigrate(&repository.Events{})
}
