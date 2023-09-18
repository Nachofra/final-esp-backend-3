package main

import (
	"github.com/Nachofra/final-esp-backend-3/cmd/api/config"
	"github.com/Nachofra/final-esp-backend-3/cmd/api/handlers/v1"
	"github.com/Nachofra/final-esp-backend-3/pkg/db/mysql"
	"github.com/Nachofra/final-esp-backend-3/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

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
	))
	if err != nil {
		panic(err)
	}

	v := validator.New(validator.WithRequiredStructEnabled())

	eng := gin.New()
	eng.Use(middleware.Logger())
	//eng.Use(middleware.Authenticate())

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	v1.Routes(eng, v1.Config{
		Log:       logger,
		DB:        database,
		Validator: v,
	})

	err = eng.Run()
	if err != nil {
		panic(err)
	}
}
