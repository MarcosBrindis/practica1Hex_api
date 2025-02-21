package controller

import (
	"context"
	"net/http"
	"practica1/src/FilmHub/application"
	"practica1/src/FilmHub/domain/model"

	"github.com/gin-gonic/gin"
)

type CreateFilmHubController struct {
	UseCase *application.CreateFilmHubUsecase
}

func NewCreateFilmHubController(useCase *application.CreateFilmHubUsecase) *CreateFilmHubController {
	return &CreateFilmHubController{UseCase: useCase}
}

func (c *CreateFilmHubController) HandleCreate(ctx *gin.Context) {
	var film model.FilmHub
	if err := ctx.ShouldBindJSON(&film); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqContext := context.Background()
	if err := c.UseCase.Execute(reqContext, &film); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "FilmHub creado exitosamente"})
}
