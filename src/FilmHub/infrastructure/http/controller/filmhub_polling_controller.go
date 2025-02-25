package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"practica1/src/FilmHub/application"

	"github.com/gin-gonic/gin"
)

// FilmHubPollingController gestiona endpoints de polling para FilmHub.
type FilmHubPollingController struct {
	GetAllFilmHubUsecase *application.GetAllFilmHubUsecase
	Updates              *chan bool
}

// Constructor
func NewFilmHubPollingController(getAllUsecase *application.GetAllFilmHubUsecase, updates *chan bool) *FilmHubPollingController {
	return &FilmHubPollingController{
		GetAllFilmHubUsecase: getAllUsecase,
		Updates:              updates,
	}
}

// HandleShortPoll responde inmediatamente con la lista actual de FilmHub.
func (c *FilmHubPollingController) HandleShortPoll(ctx *gin.Context) {
	reqCtx := context.Background()
	films, err := c.GetAllFilmHubUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, films)
}

// HandleLongPoll espera hasta 30 segundos o hasta que se reciba una notificación y luego devuelve la lista actual.
func (c *FilmHubPollingController) HandleLongPoll(ctx *gin.Context) {
	select {
	case <-*c.Updates:
		println("Se recibió una notificación de cambio")
		*c.Updates = make(chan bool, 1)
	case <-time.After(30 * time.Second):
		println("Timeout: No hubo cambios en 30 segundos")
	}

	reqCtx := context.Background()
	films, err := c.GetAllFilmHubUsecase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, films)
}

// HandleCountShortPollStreaming envía periódicamente el conteo de películas en una conexión persistente.
func (c *FilmHubPollingController) HandleCountShortPollStreaming(ctx *gin.Context) {
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

	// Definimos el tiempo total de la conexión (por ejemplo, 30 segundos)
	timeout := time.After(30 * time.Second)
	// Y el intervalo entre cada actualización (por ejemplo, cada 2 segundos)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Bucle para enviar actualizaciones hasta que se agote el timeout.
	for {
		select {
		case <-timeout:
			// Al finalizar el tiempo, terminamos la conexión.
			fmt.Fprintf(writer, "data: %s\n\n", "timeout")
			flusher.Flush()
			return
		case <-ticker.C:
			reqCtx := context.Background()
			films, err := c.GetAllFilmHubUsecase.Execute(reqCtx)
			if err != nil {
				// Si hay error, lo enviamos en el stream.
				fmt.Fprintf(writer, "data: error: %s\n\n", err.Error())
				flusher.Flush()
				continue
			}
			count := len(films)
			// Enviamos el contador en formato JSON o como string.
			fmt.Fprintf(writer, "data: {\"count\": %d}\n\n", count)
			flusher.Flush()
		}
	}
}
