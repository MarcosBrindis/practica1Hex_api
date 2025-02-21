package controller

import (
	"context"
	"net/http"
	"practica1/src/user/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	UseCase *application.DeleteUserUsecase
}

func NewDeleteUserController(useCase *application.DeleteUserUsecase) *DeleteUserController {
	return &DeleteUserController{UseCase: useCase}
}

func (c *DeleteUserController) HandleDelete(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
