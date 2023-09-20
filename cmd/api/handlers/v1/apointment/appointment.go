package apointment

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/en_validator"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
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

// Create is the handler responsible for creating a new appointment.
// @Summary Create a new appointment
// @Description Create a new appointment with JSON input
// @Tags appointment
// @Accept json
// @Produce json
// @Param request body appointment.NewAppointment true "Appointment data"
// @Success 201 {object} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
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

// GetAll is the handler responsible for retrieving all appointments.
// @Summary Get all appointments
// @Description Get a list of all appointments with optional query parameters
// @Tags appointment
// @Accept json
// @Produce json
// @Param filters query appointment.FilterAppointment false "Optional filters" default({}) Example({"dni":"12345678", "from_date":""2023-09-15 11:30:00""})
// @Success 200 {array} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment [get]
func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filters appointment.FilterAppointment

		err := ctx.ShouldBindQuery(&filters)
		if err != nil {
			msg := ""

			if strings.Contains(err.Error(), "top-level") {
				msg = ": If you are using dates via query parameters, please ensure they are wrapped in quotes."
			}

			web.Error(ctx, http.StatusBadRequest, "%s%s", err, msg)
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

// GetByID is the handler responsible for retrieving an appointment by its ID.
// @Summary Get an appointment by ID
// @Description Get an appointment by its unique ID
// @Tags appointment
// @Param id path int true "Appointment ID"
// @Accept json
// @Produce json
// @Success 200 {object} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/{id} [get]
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

// Update is the handler responsible for updating an appointment by its ID.
// @Summary Update an appointment by ID
// @Description Update an appointment with JSON input by its unique ID
// @Tags appointment
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param request body appointment.UpdateAppointment true "Updated appointment data"
// @Success 200 {object} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/{id} [put]
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

		_, err = h.service.GetByID(ctx, idInt)
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

// Patch is the handler responsible for partially updating an appointment by its ID.
// @Summary Partially update an appointment by ID
// @Description Partially update an appointment with JSON input by its unique ID
// @Tags appointment
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param request body appointment.PatchAppointment true "Partial update data"
// @Success 200 {object} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/{id} [patch]
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

// Delete is the handler responsible for deleting an appointment by its ID.
// @Summary Delete an appointment by ID
// @Description Delete an appointment by its unique ID
// @Tags appointment
// @Param id path int true "Appointment ID"
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/{id} [delete]
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
			case errors.Is(err, appointment.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err)
				return
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

// CreateByDNI is the handler responsible for creating appointments by patient DNI and dentist registration number.
// @Summary Create an appointment by patient DNI and dentist registration number
// @Description Create a new appointment with JSON input using patient DNI and dentist registration number
// @Tags appointment
// @Accept json
// @Produce json
// @Param request body appointment.NewAppointmentDNIRegistrationNumber true "Appointment data"
// @Success 201 {object} appointment.Appointment
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/dni [post]
func (h *Handler) CreateByDNI() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request appointment.NewAppointmentDNIRegistrationNumber
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
			switch {
			case errors.Is(err, patient.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}
		de, err = h.dentistService.GetByRegistrationNumber(ctx, request.DentistNumber)
		if err != nil {
			switch {
			case errors.Is(err, dentist.ErrNotFound):
				web.Error(ctx, http.StatusNotFound, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
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
