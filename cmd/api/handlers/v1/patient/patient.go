package patient

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/en_validator"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var (
	ErrInvalidID      = errors.New("invalid ID")
	ErrInternalServer = errors.New("internal server error")
)

type Handler struct {
	service   patient.Service
	validator *en_validator.Validator
}

func NewHandler(service patient.Service, validator *en_validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// Create is the handler in charge of the patient creation flow.
// Patient godoc
// @Summary patient example
// @Description Create a new patient
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient [post]
func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request patient.NewPatient

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

		p, err := h.service.Create(ctx, request)
		if err != nil {
			switch {
			case errors.Is(err, patient.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusCreated, p)
	}
}

// GetAll is the handler in charge of patient querying flow.
// Patient godoc
// @Summary patient example
// @Description Get all patients
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /patient [get]
func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := h.service.GetAll(ctx)

		web.Success(ctx, http.StatusOK, p)
	}
}

// GetByID is the handler in charge of querying patients by ID.
// Patient godoc
// @Summary patient example
// @Description Get patient by id
// @Tags patient
// @Param id path int true "patient id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [get]
func (h *Handler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		p, err := h.service.GetByID(ctx, id)
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

		web.Success(ctx, http.StatusOK, p)
	}
}

// Update is the handler in charge of patient updating flow.
// Patient godoc
// @Summary patient example
// @Description Update patient by id
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [put]
func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request patient.NewPatient

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

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		p, err := h.service.Update(ctx, request, idInt)
		if err != nil {
			switch {
			case errors.Is(err, patient.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusOK, p)
	}
}

// Patch is the handler in charge of patient updating flow.
// Patient godoc
// @Summary patient example
// @Description Patch patient by id
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [patch]
func (h *Handler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pp patient.PatchPatient
		var pa patient.Patient

		err := ctx.ShouldBindJSON(&pp)
		if err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		err = h.validator.Validate.Struct(pp)
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

		pa, err = h.service.GetByID(ctx, idInt)
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

		p, err := h.service.Patch(ctx, pa, pp)
		if err != nil {
			switch {
			case errors.Is(err, patient.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, patient.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusOK, p)
	}
}

// Delete is the handler in charge of patient deleting flow.
// Patient godoc
// @Summary patient example
// @Description Delete patient by id
// @Tags patient
// @Param id path int true "patient id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [delete]
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
			case errors.Is(err, patient.ErrConflict):
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
