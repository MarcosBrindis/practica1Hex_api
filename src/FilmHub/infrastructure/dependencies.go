package infrastructure

import (
	"fmt"
	"practica1/src/FilmHub/application"
	repo "practica1/src/FilmHub/infrastructure/database/postgres"
	"practica1/src/FilmHub/infrastructure/http/controller"
	db "practica1/src/core/database/postgres"
)

var (
	CreateFilmHubController *controller.CreateFilmHubController
	UpdateFilmHubController *controller.UpdateFilmHubController
	GetFilmHubController    *controller.GetFilmHubController
	DeleteFilmHubController *controller.DeleteFilmHubController
	GetAllFilmHubController *controller.GetAllFilmHubController
)

func InitDependencies() {
	// Inicializar la base de datos

	database, err := db.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Inicializar el repositorio de FilmHub
	filmHubRepo := repo.NewFilmHubRepository(database)

	// Crear los casos de uso
	createFilmHubUsecase := application.CreateFilmHubUsecase{Repository: filmHubRepo}
	updateFilmHubUsecase := application.UpdateFilmHubUsecase{Repository: filmHubRepo}
	getFilmHubUsecase := application.GetFilmHubUsecase{Repository: filmHubRepo}
	deleteFilmHubUsecase := application.DeleteFilmHubUsecase{Repository: filmHubRepo}
	getAllFilmHubUsecase := application.GetAllFilmHubUsecase{Repository: filmHubRepo}

	// Crear instancias de los controladores
	CreateFilmHubController = controller.NewCreateFilmHubController(&createFilmHubUsecase)
	UpdateFilmHubController = controller.NewUpdateFilmHubController(&updateFilmHubUsecase)
	GetFilmHubController = controller.NewGetFilmHubController(&getFilmHubUsecase)
	DeleteFilmHubController = controller.NewDeleteFilmHubController(&deleteFilmHubUsecase)
	GetAllFilmHubController = controller.NewGetAllFilmHubController(&getAllFilmHubUsecase)
}
