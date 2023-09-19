package dentist

import (
	"errors"
	"github.com/Nachofra/final-esp-backend-3/pkg/en_validator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidID      = errors.New("invalid ID")
	ErrInternalServer = errors.New("internal server error")
)

// Handler is a structure for dentist handler.
type Handler struct {
	service   dentist.Service
	validator *en_validator.Validator
}

// NewHandler is a function to create a handler
func NewHandler(service dentist.Service, validator *en_validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// Create is the handler in charge of the dentist creation flow.
// Dentist godoc
// @Summary dentist example
// @Description Create a new dentist
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist [post]
func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request dentist.NewDentist

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

		d, err := h.service.Create(ctx, request)
		if err != nil {
			switch {
			case errors.Is(err, dentist.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusCreated, d)
	}
}

// GetAll is the handler in charge of dentist querying flow.
// Dentist godoc
// @Summary dentist example
// @Description Get all dentists
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /dentist [get]
func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		d := h.service.GetAll(ctx)

		web.Success(ctx, http.StatusOK, d)
	}
}

// GetByID is the handler in charge of querying dentists by ID.
// Dentist godoc
// @Summary dentist example
// @Description Get dentist by id
// @Tags dentist
// @Param id path int true "dentist id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [get]
func (h *Handler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", ErrInvalidID)
			return
		}

		d, err := h.service.GetByID(ctx, id)
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

		web.Success(ctx, http.StatusOK, d)
	}
}

// Update is the handler in charge of dentist updating flow.
// Dentist godoc
// @Summary dentist example
// @Description Update dentist by id
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [put]
func (h *Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request dentist.UpdateDentist

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

		d, err := h.service.Update(ctx, request, idInt)
		if err != nil {
			switch {
			case errors.Is(err, dentist.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusOK, d)
	}
}

// Delete is the handler in charge of dentist deleting flow.
// Dentist godoc
// @Summary dentist example
// @Description Delete dentist by id
// @Tags dentist
// @Param id path int true "dentist id"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [delete]
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
			case errors.Is(err, dentist.ErrConflict):
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

// Patch is the handler in charge of appointment patching flow.
// Dentist godoc
// @Summary dentist example
// @Description Patch dentist by id
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [patch]
func (h *Handler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var de dentist.Dentist
		var pd dentist.PatchDentist

		errBind := ctx.ShouldBindJSON(&pd)
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

		de, err = h.service.GetByID(ctx, idInt)
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

		d, err := h.service.Patch(ctx, de, pd)
		if err != nil {
			switch {
			case errors.Is(err, dentist.ErrAlreadyExists):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrConflict):
				web.Error(ctx, http.StatusConflict, "%s", err)
				return
			case errors.Is(err, dentist.ErrValueExceeded):
				web.Error(ctx, http.StatusUnprocessableEntity, "%s", err)
				return
			default:
				web.Error(ctx, http.StatusInternalServerError, "%s", ErrInternalServer)
				return
			}
		}

		web.Success(ctx, http.StatusOK, d)
	}
}
