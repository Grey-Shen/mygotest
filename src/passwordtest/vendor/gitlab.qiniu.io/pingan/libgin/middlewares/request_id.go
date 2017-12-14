package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AssignRequestID(contextKey, header string) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := uuid.NewV4().String()
		context.Set(contextKey, id)
		context.Writer.Header().Add(header, id)
		context.Next()
	}
}
