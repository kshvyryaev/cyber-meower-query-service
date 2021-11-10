package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http/response"
	"go.uber.org/zap"
)

type ErrorHandlerMiddleware struct {
	logger *zap.Logger
}

func ProvideErrorHandlerMiddleware(logger *zap.Logger) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{
		logger: logger,
	}
}

func (handler *ErrorHandlerMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors[0].Err
			context.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})

			handler.logger.Error("error happend: " + err.Error())
		}
	}
}
