package rest_api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			map[string]interface{}{
				"timestamp": time.Now().Unix(),
				"version":   os.Getenv("VERSION"),
			},
		)
	}
}
