package apointment

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/time"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	ErrInvalidID           = errors.New("invalid ID")
	ErrInternalServer      = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity: the JSON provided does not conform to the expected entity structure, please review it and try again")
)

type Handler struct {
	service        appointment.Service
	patientService patient.Service
	dentistService dentist.Service
}

func NewHandler(service appointment.Service, patientService patient.Service, dentistService dentist.Service) *Handler {
	return &Handler{
		service:        service,
		patientService: patientService,
		dentistService: dentistService,
	}
}

// Create is the handler in charge of the appointment creation flow.
// Appointment godoc
// @Summary appointment example
// @Description Create a new appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment [post]
func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request appointment.NewAppointment

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}

		app, err := h.service.Create(ctx, request)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}
		web.Success(ctx, http.StatusCreated, app)
	}
}

// GetAll is the handler in charge of appointment querying flow.
// Appointment godoc
// @Summary appointment example
// @Description Get all appointments
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /appointment [get]
func (h *Handler) GetByDNI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dni, err := strconv.Atoi(ctx.Query("dni"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID.Error())
			return
		}

		appointments := h.service.GetAll(ctx, appointment.FilterAppointment{DNI: dni})

		web.Success(ctx, http.StatusOK, appointments)
	}
}

// GetByID is the handler in charge of querying appointments by ID.
// Appointment godoc
// @Summary appointment example
// @Description Get appointment by id
// @Tags appointment
// @Param id path int true "appointment id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [get]
func (h *Handler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID.Error())
			return
		}

		app, err := h.service.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}

		web.Success(ctx, http.StatusOK, app)
	}
}

// Update is the handler in charge of appointment updating flow.
// Appointment godoc
// @Summary appointment example
// @Description Update appointment by id
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [put]
func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request appointment.Appointment
		var ua appointment.UpdateAppointment

		errBind := ctx.ShouldBindJSON(&ua)
		if errBind != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID.Error())
			return
		}

		request.ID = idInt

		app, err := h.service.Update(ctx, request, ua)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}
		web.Success(ctx, http.StatusOK, app)
	}
}

// PatchUpdate is the handler in charge of appointment updating flow.
// Appointment godoc
// @Summary appointment example
// @Description PatchUpdate appointment by id
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [put]
func (h *Handler) PatchUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request appointment.Appointment
		var pa appointment.PatchAppointment

		errBind := ctx.ShouldBindJSON(&pa)
		if errBind != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID.Error())
			return
		}

		request, _ = h.service.GetByID(ctx, idInt)

		app, err := h.service.Patch(ctx, request, pa)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}
		web.Success(ctx, http.StatusOK, app)
	}
}

// Delete is the handler in charge of appointment deleting flow.
// Appointment godoc
// @Summary appointment example
// @Description Delete appointment by id
// @Tags appointment
// @Param id path int true "appointment id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [delete]
func (h *Handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID.Error())
			return
		}

		err = h.service.Delete(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}
		web.Success(ctx, http.StatusNoContent, nil)
	}
}

// CreateByDNI is the handler in charge of creating appointments by DNI and RegistrationNumber.
// Appointment by DNI godoc
// @Summary appointment example
// @Description Create a new appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/dni [post]
func (h *Handler) CreateByDNI() gin.HandlerFunc {
	// TODO crear metodo para crear turno con dni de paciente
	return func(ctx *gin.Context) {

		type NewCreate struct {
			PatientDNI    int       `json:"patient_dni"`
			DentistNumber int       `json:"dentist_number"`
			Date          time.Time `json:"date"`
			Description   string    `json:"description"`
		}

		var request NewCreate
		var app appointment.NewAppointment
		var pa patient.Patient
		var de dentist.Dentist

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}

		pa, err = h.patientService.GetByDNI(ctx, request.PatientDNI)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}
		de, err = h.dentistService.GetByRegistrationNumber(ctx, request.DentistNumber)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", ErrUnprocessableEntity.Error())
			return
		}

		app.PatientID = pa.ID
		app.DentistID = de.ID
		app.Date = request.Date
		app.Description = request.Description

		newApp, err := h.service.Create(ctx, app)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err.Error())
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err.Error())
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer.Error())
				return
			}
		}
		web.Success(ctx, http.StatusCreated, newApp)
	}
}
