package route

import (
	"real-estate-api/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func FileStorageRoutes(engine *gin.Engine, fileHandler *handler.FileStorageHandler) {
	engine.POST("/real-estate/image", fileHandler.Upload)

}
