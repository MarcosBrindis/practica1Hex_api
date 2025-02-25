package router

import (
	"practica1/src/FilmHub/infrastructure/http/controller"
	"practica1/src/core/middleware"

	"github.com/gin-gonic/gin"
)

func SetupFilmHubRoutes(r *gin.Engine,
	create *controller.CreateFilmHubController,
	update *controller.UpdateFilmHubController,
	get *controller.GetFilmHubController,
	delete *controller.DeleteFilmHubController,
	getAll *controller.GetAllFilmHubController,
	filmhubPolling *controller.FilmHubPollingController,
	updates *chan bool,
) {

	filmhubGroup := r.Group("/filmhub")
	{
		filmhubGroup.POST("/", middleware.NotifyUpdatesMiddleware(updates), create.HandleCreate)
		filmhubGroup.PUT("/:id", middleware.NotifyUpdatesMiddleware(updates), update.HandleUpdate)
		filmhubGroup.DELETE("/:id", middleware.NotifyUpdatesMiddleware(updates), delete.HandleDelete)
		filmhubGroup.GET("/:id", get.HandleGet)
		filmhubGroup.GET("/", getAll.HandleGetAll)

		// Endpoints para polling
		filmhubGroup.GET("/shortpoll", filmhubPolling.HandleShortPoll)
		filmhubGroup.GET("/longpoll", filmhubPolling.HandleLongPoll)
		filmhubGroup.GET("/poll/count", filmhubPolling.HandleCountShortPollStreaming)
	}
}
