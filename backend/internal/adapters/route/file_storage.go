package route

import (
	"real-state-api/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func FileStorageRoutes(engine *gin.Engine, fileHandler *handler.FileStorageHandler) {
	engine.POST("/realestate/image", fileHandler.Upload)

}
