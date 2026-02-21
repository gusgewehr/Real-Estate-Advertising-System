package route

import (
	"real-state-api/internal/adapters/handler"

	"github.com/gin-gonic/gin"
)

func ExchangeRateRoutes(engine *gin.Engine, exchangeRateHandler *handler.ExchangeRateHandler) {
	engine.POST("/exchange-rate", exchangeRateHandler.CreateExchangeRate)
	engine.GET("/exchange-rate/latest", exchangeRateHandler.GetLatestExchangeRate)
}
