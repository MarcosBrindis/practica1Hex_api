package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotifyUpdatesMiddleware es un middleware que notifica cambios al canal Updates.
func NotifyUpdatesMiddleware(updates *chan bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ejecutar la siguiente función en la cadena de middlewares
		ctx.Next()

		// Verificar si la solicitud fue exitosa (código 2xx)
		if ctx.Writer.Status() >= http.StatusOK && ctx.Writer.Status() < http.StatusMultipleChoices {
			// Notificar que hubo un cambio
			select {
			case *updates <- true: // Enviar señal al canal
				println("Notificación enviada al canal Updates")
			default: // Evitar bloqueo si el canal está lleno
				println("Canal Updates lleno, no se pudo enviar notificación")
			}
		}
	}
}
