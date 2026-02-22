package route

import (
	"real-estate-api/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func RealEstateRoutes(engine *gin.Engine, realEstateHandler *handler.RealEstateHandler) {

	engine.POST("/real-estate", realEstateHandler.CreateRealEstate)
	engine.GET("/real-estate", realEstateHandler.ListRealEstates)

}
