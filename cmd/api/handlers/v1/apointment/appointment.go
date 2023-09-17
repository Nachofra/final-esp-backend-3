package apointment

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	ErrInternalServer = errors.New("internal server error")
)

type Handler struct {
	service appointment.Service
}

func NewHandler(service appointment.Service) *Handler {
	return &Handler{
		service: service,
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

		err := ctx.Bind(&request)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		app, err := h.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": app,
		})

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

		// TODO agregar filtro para buscar por dni de paciente
		appointment := h.service.GetAll(ctx, appointment.FilterAppointment{})

		web.Success(ctx, http.StatusOK, gin.H{
			"data": appointment,
		})
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
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		app, err := h.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": app,
		})
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

		errBind := ctx.ShouldBind(&ua)
		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		request.ID = idInt

		app, err := h.service.Update(ctx, request, ua)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": app,
		})

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
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		err = h.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "Appointment deleted",
		})
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
// @Router /appointment/DNI [post]
func (h *Handler) CreateByDNI() gin.HandlerFunc {
	// TODO crear metodo para crear turno con dni de paciente
	return func(ctx *gin.Context) {

		var request appointment.NewAppointment

		err := ctx.Bind(&request)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		app, err := h.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": app,
		})

	}
}

// TODO crear metodo PATCH por algun campo de turno (por ej editar fecha)
