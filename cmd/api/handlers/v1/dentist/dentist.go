package dentist

import (
	"net/http"
	"strconv"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/pkg/web"
	"github.com/gin-gonic/gin"
)

// Here we will implement dentist handlers

type Handler struct {
	// Here we put the service
	service dentist.Service
}

func NewHandler(service dentist.Service) *Handler {
	return &Handler{
		// Receive service via method param
		service: service,
	}
}

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

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		product, err := h.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": product,
		})

	}
}

// Dentist godoc
// @Summary denist example
// @Description Get all dentists
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /dentist [get]
func (h *Handler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentist := h.service.GetAll(ctx)

		web.Success(ctx, http.StatusOK, gin.H{
			"data": dentist,
		})
	}
}

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
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		dentist, err := h.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": dentist,
		})
	}
}

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

		var request dentist.NewDentist
		errBind := ctx.Bind(&request)

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

		dentist, err := h.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": dentist,
		})

	}
}

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
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		err = h.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "dentist deleted",
		})
	}
}
