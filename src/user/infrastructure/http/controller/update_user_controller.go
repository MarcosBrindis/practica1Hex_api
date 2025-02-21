package controller

import (
	"context"
	"net/http"
	"practica1/src/user/application"
	"practica1/src/user/domain/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	UseCase *application.UpdateUserUsecase
}

func NewUpdateUserController(useCase *application.UpdateUserUsecase) *UpdateUserController {
	return &UpdateUserController{UseCase: useCase}
}

func (c *UpdateUserController) HandleUpdate(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	if err := c.UseCase.Execute(reqContext, id, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}
