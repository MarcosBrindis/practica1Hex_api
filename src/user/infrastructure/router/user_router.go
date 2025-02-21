package router

import (
	"practica1/src/user/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine,
	create *controller.CreateUserController,
	update *controller.UpdateUserController,
	get *controller.GetUserController,
	delete *controller.DeleteUserController,
	getAll *controller.GetAllUsersController) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/", create.HandleCreate)
		userGroup.PUT("/:id", update.HandleUpdate)
		userGroup.GET("/:id", get.HandleGet)
		userGroup.DELETE("/:id", delete.HandleDelete)
		userGroup.GET("/", getAll.HandleGetAll)
	}
}
