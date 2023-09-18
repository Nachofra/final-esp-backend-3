package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

// Config has all the dependencies and requirements to initialize handlers.
type Config struct {
	Log       *log.Logger
	DB        *sql.DB
	Validator *validator.Validate
}

// Routes sets all the version 1 routes.
func Routes(eng *gin.Engine, cfg Config) {
	cfg.Log.Println("configuring v1 routes")

	const prefix = "/v1"
	//v1 := eng.Group(prefix)

	// dentistService := dentist.NewDentist(mysql.New(cfg.DB))

	{
		//dentistHandler := dentist.NewHandler(dentistService)
		//dentist := v1.Group("dentist")
		//
		//dentist.GET()
		//etc...
	}
}
