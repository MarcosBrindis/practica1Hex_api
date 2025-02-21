package controller

import (
	"context"
	"net/http"
	"practica1/src/FilmHub/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetFilmHubController struct {
	UseCase *application.GetFilmHubUsecase
}

func NewGetFilmHubController(useCase *application.GetFilmHubUsecase) *GetFilmHubController {
	return &GetFilmHubController{UseCase: useCase}
}

func (c *GetFilmHubController) HandleGet(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	reqContext := context.Background()
	filmhub, err := c.UseCase.Execute(reqContext, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "FilmHub no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, filmhub)
}
