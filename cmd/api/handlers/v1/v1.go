package v1

import (
	"database/sql"
	handlerAppointment "github.com/Nachofra/final-esp-backend-3/cmd/api/handlers/v1/apointment"
	handlerDentist "github.com/Nachofra/final-esp-backend-3/cmd/api/handlers/v1/dentist"
	handlerPatient "github.com/Nachofra/final-esp-backend-3/cmd/api/handlers/v1/patient"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	mysqlAppointment "github.com/Nachofra/final-esp-backend-3/internal/domain/appointment/stores/mysql"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	mysqlDentist "github.com/Nachofra/final-esp-backend-3/internal/domain/dentist/stores/mysql"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	mysqlPatient "github.com/Nachofra/final-esp-backend-3/internal/domain/patient/stores/mysql"
	"github.com/Nachofra/final-esp-backend-3/pkg/middleware"
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
	v1 := eng.Group(prefix)

	repoDentist := mysqlDentist.New(cfg.DB)
	dentistService := dentist.NewService(repoDentist)

	repoPatient := mysqlPatient.NewStore(cfg.DB)
	patientService := patient.NewService(repoPatient)

	repoAppointment := mysqlAppointment.NewStore(cfg.DB)
	appointmentService := appointment.NewService(repoAppointment)

	dentistHandler := handlerDentist.NewHandler(dentistService)
	d := v1.Group("/dentist")
	{
		d.GET("/:id", dentistHandler.GetByID())
		d.GET("", dentistHandler.GetAll())
		d.POST("/", middleware.Authenticate(), dentistHandler.Create())
		d.PUT("/:id", middleware.Authenticate(), dentistHandler.Update())
		d.PATCH("/", middleware.Authenticate(), dentistHandler.Patch())
		d.DELETE("/:id", middleware.Authenticate(), dentistHandler.Delete())
	}

	patientHandler := handlerPatient.NewHandler(*patientService)
	p := v1.Group("/patient")
	{
		p.GET("/:id", patientHandler.GetByID())
		p.GET("/", patientHandler.GetAll())
		p.POST("/", middleware.Authenticate(), patientHandler.Create())
		p.PUT("/:id", middleware.Authenticate(), patientHandler.Update())
		p.PATCH("/:id", middleware.Authenticate(), patientHandler.PatchUpdate())
		p.DELETE("/:id", middleware.Authenticate(), patientHandler.Delete())
	}

	appointmentHandler := handlerAppointment.NewHandler(*appointmentService, *patientService, dentistService)
	a := v1.Group("/appointment")
	{
		a.GET("/:id", appointmentHandler.GetByID())
		a.GET("/", appointmentHandler.GetAll())
		a.POST("/", middleware.Authenticate(), appointmentHandler.Create())
		a.POST("/dni", middleware.Authenticate(), appointmentHandler.CreateByDNI())
		a.PUT("/:id", middleware.Authenticate(), appointmentHandler.Update())
		a.PATCH("/:id", middleware.Authenticate(), appointmentHandler.PatchUpdate())
		a.DELETE("/:id", middleware.Authenticate(), appointmentHandler.Delete())
	}

}
