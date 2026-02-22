package handler

import (
	"net/http"
	"real-estate-api/internal/application/port"
	"real-estate-api/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExchangeRateHandler struct {
	logger *zap.SugaredLogger
	port   port.ExchangeRateInput
}

func NewExchangeRateHandler(port port.ExchangeRateInput, logger *zap.SugaredLogger) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		logger: logger,
		port:   port,
	}
}

// Post
// @Summary Create a new exchange rate
// @Schemes
// @Description Create a new exchange rate
// @Tags Exchange Rate
// @Accept json
// @Produce json
// @Param exchangeRate body domain.ExchangeRateInput true "Exchange rate object"
// @Success 201 {object} domain.NoBodyResponse "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"
// @Router /exchange-rate [post]
func (h *ExchangeRateHandler) CreateExchangeRate(c *gin.Context) {
	var rate *domain.ExchangeRateInput

	err := c.ShouldBindJSON(&rate)
	if err != nil {
		h.logger.Errorw("Request did not bind to json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request did not bind to json"})
		return
	}

	err = h.port.Create(rate)
	if err != nil {
		h.logger.Errorw("Error creating exchange rate", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating exchange rate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Exchange rate created successfully"})

}

// Get
// @Summary Get latest exchange rate
// @Schemes
// @Description Get latest exchange rate
// @Tags Exchange Rate
// @Accept json
// @Produce json
// @Success 200 {object} domain.ExchangeRateInput "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"// @Router /exchange-rate/latest [get]
func (h *ExchangeRateHandler) GetLatestExchangeRate(c *gin.Context) {

	rate, err := h.port.GetLatest()
	if err != nil {
		h.logger.Errorw("Error getting latest exchange rate", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting latest exchange rate"})
		return
	}

	c.JSON(http.StatusOK, rate)

}
