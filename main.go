package main

import (
	filmhubInfra "practica1/src/FilmHub/infrastructure"
	filmhubRouter "practica1/src/FilmHub/infrastructure/router"
	userInfra "practica1/src/user/infrastructure"
	userRouter "practica1/src/user/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar dependencias
	userInfra.InitDependencies()
	filmhubInfra.InitDependencies()

	// Crear instancia de Gin
	r := gin.Default()

	// Configurar rutas con los controladores globales
	userRouter.SetupUserRoutes(r, userInfra.CreateUserController, userInfra.UpdateUserController, userInfra.GetUserController, userInfra.DeleteUserController, userInfra.GetAllUsersController, userInfra.UserPollingController, &userInfra.Updates)
	filmhubRouter.SetupFilmHubRoutes(r, filmhubInfra.CreateFilmHubController, filmhubInfra.UpdateFilmHubController, filmhubInfra.GetFilmHubController, filmhubInfra.DeleteFilmHubController, filmhubInfra.GetAllFilmHubController, filmhubInfra.FilmHubPollingController, &filmhubInfra.Updates)

	// Iniciar el servidor
	r.Run(":8080")
}
