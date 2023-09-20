package main

import (
	"github.com/Nachofra/final-esp-backend-3/cmd/api/config"
	"github.com/Nachofra/final-esp-backend-3/cmd/api/handlers/v1"
	"github.com/Nachofra/final-esp-backend-3/pkg/db/mysql"
	"github.com/Nachofra/final-esp-backend-3/pkg/en_validator"
	"github.com/Nachofra/final-esp-backend-3/pkg/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// @title Final Backend Specialization 3
// @version 1.0
// @description This API handles patients, appointments and dentists.
// @BasePath /v1
func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	database, err := mysql.Open(mysql.New(
		mysql.WithUsername(cfg.DBUser),
		mysql.WithPassword(cfg.DBPassword),
		mysql.WithHost(cfg.DBHost+":"+cfg.DBPort),
		mysql.WithName(cfg.DBSchema),
		mysql.WithCharset(cfg.DBCharset),
		mysql.WithParseTime(cfg.DBParseTime),
	))
	if err != nil {
		panic(err)
	}

	eng := gin.New()
	eng.Use(middleware.Logger())

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	validator := en_validator.Get()

	v1.Routes(eng, v1.Config{
		Log:       logger,
		DB:        database,
		Validator: validator,
		Env:       cfg,
	})

	err = eng.Run(cfg.Host + ":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
