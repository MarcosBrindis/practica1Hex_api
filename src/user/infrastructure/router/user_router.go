package router

import (
	"practica1/src/core/middleware"
	"practica1/src/user/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine,
	create *controller.CreateUserController,
	update *controller.UpdateUserController,
	get *controller.GetUserController,
	delete *controller.DeleteUserController,
	getAll *controller.GetAllUsersController,
	userPolling *controller.UserPollingController,
	updates *chan bool,
) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/", middleware.NotifyUpdatesMiddleware(updates), create.HandleCreate)
		userGroup.PUT("/:id", middleware.NotifyUpdatesMiddleware(updates), update.HandleUpdate)
		userGroup.DELETE("/:id", middleware.NotifyUpdatesMiddleware(updates), delete.HandleDelete)
		userGroup.GET("/:id", get.HandleGet)
		userGroup.GET("/", getAll.HandleGetAll)

		// Endpoints para polling
		userGroup.GET("/shortpoll", userPolling.HandleShortPoll)
		userGroup.GET("/longpoll", userPolling.HandleLongPoll)
		userGroup.GET("/poll/count", userPolling.HandleCountShortPollStreaming)
	}
}
