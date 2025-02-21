package controller

import (
	"context"
	"net/http"
	"practica1/src/FilmHub/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteFilmHubController struct {
	UseCase *application.DeleteFilmHubUsecase
}

func NewDeleteFilmHubController(useCase *application.DeleteFilmHubUsecase) *DeleteFilmHubController {
	return &DeleteFilmHubController{UseCase: useCase}
}

func (c *DeleteFilmHubController) HandleDelete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "FilmHub eliminado exitosamente"})
}
