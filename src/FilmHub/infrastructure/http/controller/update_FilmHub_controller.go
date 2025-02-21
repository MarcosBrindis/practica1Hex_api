package controller

import (
	"context"
	"net/http"
	"practica1/src/FilmHub/application"
	"practica1/src/FilmHub/domain/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateFilmHubController struct {
	UseCase *application.UpdateFilmHubUsecase
}

func NewUpdateFilmHubController(useCase *application.UpdateFilmHubUsecase) *UpdateFilmHubController {
	return &UpdateFilmHubController{UseCase: useCase}
}

func (c *UpdateFilmHubController) HandleUpdate(ctx *gin.Context) {
	var film model.FilmHub
	if err := ctx.ShouldBindJSON(&film); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, id, &film); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "FilmHub actualizado exitosamente"})
}
