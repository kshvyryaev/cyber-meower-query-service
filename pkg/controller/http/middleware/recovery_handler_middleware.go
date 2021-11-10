package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http/response"
	"go.uber.org/zap"
)

type RecoveryHandlerMiddleware struct {
	logger *zap.Logger
}

func ProvideRecoveryHandlerMiddleware(logger *zap.Logger) *RecoveryHandlerMiddleware {
	return &RecoveryHandlerMiddleware{
		logger: logger,
	}
}

func (handler *RecoveryHandlerMiddleware) Handle() gin.HandlerFunc {
	return gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		err, ok := recovered.(string)

		if ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: err})
		} else {
			context.AbortWithStatus(http.StatusInternalServerError)
		}

		handler.logger.Error("panic happend: " + err)
	})
}
