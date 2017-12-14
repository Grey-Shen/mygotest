package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Header(name, value string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add(name, value)
	}
}
