package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"practica1/src/user/application"

	"github.com/gin-gonic/gin"
)

// UserPollingController gestiona endpoints de polling para usuarios.
type UserPollingController struct {
	GetAllUsersUsecase *application.GetAllUsersUsecase
	Updates            *chan bool
}

// Constructor
func NewUserPollingController(getAllUsecase *application.GetAllUsersUsecase, updates *chan bool) *UserPollingController {
	return &UserPollingController{
		GetAllUsersUsecase: getAllUsecase,
		Updates:            updates,
	}
}

// HandleShortPoll responde inmediatamente con el estado actual de los usuarios.
func (c *UserPollingController) HandleShortPoll(ctx *gin.Context) {
	reqCtx := context.Background()
	users, err := c.GetAllUsersUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// HandleLongPoll espera hasta 30 segundos o hasta que se reciba una notificación de cambio, y luego devuelve la lista actual.
func (c *UserPollingController) HandleLongPoll(ctx *gin.Context) {
	select {
	case <-*c.Updates:
		println("Se recibió una notificación de cambio")
		*c.Updates = make(chan bool, 1)
	case <-time.After(30 * time.Second):
		println("Timeout: No hubo cambios en 30 segundos")
	}

	reqCtx := context.Background()
	users, err := c.GetAllUsersUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// HandleCountShortPollStreaming envía periódicamente el conteo de usuarios en una conexión persistente.
func (c *UserPollingController) HandleCountShortPollStreaming(ctx *gin.Context) {
	// Configuramos los headers para streaming con chunked encoding.
	writer := ctx.Writer
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")
	// Verificar que el writer soporte flushing.
	flusher, ok := writer.(http.Flusher)
	if !ok {
		ctx.String(http.StatusInternalServerError, "Streaming no soportado")
		return
	}

	// Definimos un timeout total de 30 segundos y un ticker que dispare cada 2 segundos.
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			// Al finalizar el tiempo, terminamos la conexión.
			fmt.Fprintf(writer, "data: %s\n\n", "timeout")
			flusher.Flush()
			return
		case <-ticker.C:
			reqCtx := context.Background()
			users, err := c.GetAllUsersUsecase.Execute(reqCtx)
			if err != nil {
				// Si hay error, lo enviamos en el stream.
				fmt.Fprintf(writer, "data: error: %s\n\n", err.Error())
				flusher.Flush()
				continue
			}
			count := len(users)
			// Enviamos el contador en formato JSON.
			fmt.Fprintf(writer, "data: {\"count\": %d}\n\n", count)
			flusher.Flush()
		}
	}
}
