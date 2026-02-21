package route

import (
	"real-state-api/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func ZipCodeRoutes(engine *gin.Engine, zipcodeHandler *handler.ZipCodeHandler) {
	engine.GET("/zipcode/:zipCode", zipcodeHandler.GetZipCode)
}
