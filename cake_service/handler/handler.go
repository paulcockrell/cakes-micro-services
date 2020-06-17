package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulcockrell/waracle-cake-service/repository"
)

type Handler struct {
	Repository repository.Repository
}

// Create - Create cake action
func (h *Handler) Create(ctx *gin.Context) {
	var cake repository.Cake
	err := ctx.BindJSON(&cake)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	err = h.Repository.Create(ctx, &cake)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cake)
}

// GetAll - Get all cakes
func (h *Handler) GetAll(ctx *gin.Context) {
	cakes, err := h.Repository.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cakes)
}

// Get - Get cake by ID
func (h *Handler) Get(ctx *gin.Context) {
	cake, err := h.Repository.Get(ctx, ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, cake)
}

// Update - Update cake
func (h *Handler) Update(ctx *gin.Context) {
	var cake repository.Cake
	ctx.BindJSON(&cake)

	updatedCake, err := h.Repository.Update(ctx, ctx.Param("id"), &cake)
	if err != nil {
		ctx.Error(err)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, updatedCake)
}

// Delete - Delete a cake
func (h *Handler) Delete(ctx *gin.Context) {
	err := h.Repository.Delete(ctx, ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
