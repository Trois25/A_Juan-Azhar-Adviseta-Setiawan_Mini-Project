package main

import (
	"event_ticket/app/configs"
	"event_ticket/app/database"
	"event_ticket/app/migration"

	// "event_ticket/app/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(cfg)
	migration.InitMigrationMysql(db)

	e := echo.New()

	// router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
