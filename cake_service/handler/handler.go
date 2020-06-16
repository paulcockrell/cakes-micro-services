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
	ctx.BindJSON(&cake)

	if err := h.Repository.Create(ctx, &cake); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, cake)
}

func (h *Handler) GetAll(ctx *gin.Context) {
	var cakes []*repository.Cake
	cakes, err := h.Repository.GetAll(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, cakes)
}
