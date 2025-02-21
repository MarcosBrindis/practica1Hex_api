package controller

import (
	"context"
	"net/http"
	"practica1/src/FilmHub/application"

	"github.com/gin-gonic/gin"
)

type GetAllFilmHubController struct {
	UseCase *application.GetAllFilmHubUsecase
}

func NewGetAllFilmHubController(useCase *application.GetAllFilmHubUsecase) *GetAllFilmHubController {
	return &GetAllFilmHubController{UseCase: useCase}
}

func (c *GetAllFilmHubController) HandleGetAll(ctx *gin.Context) {
	reqContext := context.Background()
	filmhubs, err := c.UseCase.Execute(reqContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, filmhubs)
}
