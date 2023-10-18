package database

import (
	"event_ticket/app/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *configs.AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
