package infrastructure

import (
	"fmt"
	db "practica1/src/core/database/postgres"
	"practica1/src/user/application"
	repo "practica1/src/user/infrastructure/database/postgres"
	"practica1/src/user/infrastructure/http/controller"
	bcryptEnc "practica1/src/user/infrastructure/service"
)

var (
	CreateUserController  *controller.CreateUserController
	UpdateUserController  *controller.UpdateUserController
	GetUserController     *controller.GetUserController
	DeleteUserController  *controller.DeleteUserController
	GetAllUsersController *controller.GetAllUsersController
	UserPollingController *controller.UserPollingController
	Updates               chan bool
)

func InitDependencies() {
	// Inicializar la base de datos
	database, err := db.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Inicializar el repositorio de usuario
	userRepo := repo.NewUserRepository(database)

	// Configurar bcrypt
	encrypter := bcryptEnc.NewBcryptEncrypter(10)

	// Crear los casos de uso
	createUserUsecase := application.CreateUserUsecase{Repository: userRepo, Encrypter: encrypter}
	updateUserUsecase := application.UpdateUserUsecase{Repository: userRepo}
	getUserUsecase := application.GetUserUsecase{Repository: userRepo}
	deleteUserUsecase := application.DeleteUserUsecase{Repository: userRepo}
	getAllUsersUsecase := application.GetAllUsersUsecase{Repository: userRepo}

	// Crear el canal de notificaciones
	Updates = make(chan bool, 1)

	// Crear instancias de los controladores
	CreateUserController = controller.NewCreateUserController(&createUserUsecase)
	UpdateUserController = controller.NewUpdateUserController(&updateUserUsecase)
	GetUserController = controller.NewGetUserController(&getUserUsecase)
	DeleteUserController = controller.NewDeleteUserController(&deleteUserUsecase)
	GetAllUsersController = controller.NewGetAllUsersController(&getAllUsersUsecase)

	// Crear controlador de polling
	UserPollingController = controller.NewUserPollingController(&getAllUsersUsecase, &Updates)
}
