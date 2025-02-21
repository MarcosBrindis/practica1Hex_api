package router

import (
	"practica1/src/FilmHub/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupFilmHubRoutes(r *gin.Engine,
	create *controller.CreateFilmHubController,
	update *controller.UpdateFilmHubController,
	get *controller.GetFilmHubController,
	delete *controller.DeleteFilmHubController,
	getAll *controller.GetAllFilmHubController) {

	filmhubGroup := r.Group("/filmhub")
	{
		filmhubGroup.POST("/", create.HandleCreate)
		filmhubGroup.PUT("/:id", update.HandleUpdate)
		filmhubGroup.GET("/:id", get.HandleGet)
		filmhubGroup.DELETE("/:id", delete.HandleDelete)
		filmhubGroup.GET("/", getAll.HandleGetAll)
	}
}
