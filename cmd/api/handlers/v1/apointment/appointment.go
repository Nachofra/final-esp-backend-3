package apointment

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/custom_time"
	"github.com/Nachofra/final-esp-backend-3/pkg/en_validator"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var (
	ErrInvalidID      = errors.New("invalid request, please check input types")
	ErrInternalServer = errors.New("internal server error")
)

// Handler is a structure for appointment handler
type Handler struct {
	service        appointment.Service
	patientService patient.Service
	dentistService dentist.Service
	validator      *en_validator.Validator
}

// NewHandler is a function to create a handler
func NewHandler(service appointment.Service, patientService patient.Service, dentistService dentist.Service, validator *en_validator.Validator) *Handler {
	return &Handler{
		service:        service,
		patientService: patientService,
		dentistService: dentistService,
		validator:      validator,
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
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		err = h.validator.Validate.Struct(request)
		if err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			msg := h.validator.Translate(validationErrors)

			web.Error(ctx, http.StatusUnprocessableEntity, "%v", msg)
			return
		}

		app, err := h.service.Create(ctx, request)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
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
func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filters appointment.FilterAppointment

		err := ctx.ShouldBindQuery(&filters)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", err)
			return
		}

		err = h.validator.Validate.Struct(filters)
		if err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			msg := h.validator.Translate(validationErrors)

			web.Error(ctx, http.StatusUnprocessableEntity, "%v", msg)
			return
		}

		appointments := h.service.GetAll(ctx, filters)

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
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		app, err := h.service.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
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

		var ua appointment.UpdateAppointment

		err := ctx.ShouldBindJSON(&ua)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		err = h.validator.Validate.Struct(ua)
		if err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			msg := h.validator.Translate(validationErrors)

			web.Error(ctx, http.StatusUnprocessableEntity, "%v", msg)
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		app, err := h.service.Update(ctx, idInt, ua)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}
		web.Success(ctx, http.StatusOK, app)
	}
}

// Patch is the handler in charge of appointment updating flow.
// Appointment godoc
// @Summary appointment example
// @Description Patch appointment by id
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [patch]
func (h *Handler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pa appointment.PatchAppointment
		var app appointment.Appointment

		errBind := ctx.ShouldBindJSON(&pa)
		if errBind != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", errBind)
			return
		}

		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		app, err = h.service.GetByID(ctx, idInt)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		a, err := h.service.Patch(ctx, app, pa)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}
		web.Success(ctx, http.StatusOK, a)
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
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		err = h.service.Delete(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
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
	return func(ctx *gin.Context) {

		type NewCreate struct {
			PatientDNI    int              `json:"patient_dni" validate:"min=10000000,max=999999999"`
			DentistNumber int              `json:"dentist_number"`
			Date          custom_time.Time `json:"date"`
			Description   string           `json:"description"`
		}

		var request NewCreate
		var app appointment.NewAppointment
		var pa patient.Patient
		var de dentist.Dentist

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		err = h.validator.Validate.Struct(request)
		if err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			msg := h.validator.Translate(validationErrors)

			web.Error(ctx, http.StatusUnprocessableEntity, "%v", msg)
			return
		}

		pa, err = h.patientService.GetByDNI(ctx, request.PatientDNI)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
			return
		}
		de, err = h.dentistService.GetByRegistrationNumber(ctx, request.DentistNumber)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
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
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, appointment.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}
		web.Success(ctx, http.StatusCreated, newApp)
	}
}
