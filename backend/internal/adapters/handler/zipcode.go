package handler

import (
	"errors"
	"net/http"
	"real-estate-api/internal/application/port"
	"real-estate-api/internal/domain"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ZipCodeHandler struct {
	port port.ZipCodeInputPort
}

func NewZipCodeHandler(port port.ZipCodeInputPort) *ZipCodeHandler {
	return &ZipCodeHandler{port}
}

// Get
// @Summary Get zipcode information
// @Schemes
// @Description Get zipcode information
// @Tags Zipcode
// @Accept json
// @Produce json
// @Param zipCode path string true "Zipcode"
// @Success 200 {object} domain.Address "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"
// @Router /zipcode/{zipCode} [get]
func (h *ZipCodeHandler) GetZipCode(c *gin.Context) {
	zipcode := c.Param("zipCode")

	zipcode = regexp.MustCompile(`[^0-9]`).ReplaceAllString(zipcode, "")

	if zipcode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Zipcode is required"})
		return
	}

	address, err := h.port.GetZipCode(zipcode)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrConnection):
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Service unavailable"})
		case errors.Is(err, domain.ThirdPartyErr):
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Third-party service error"})
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Zipcode not found"})
		case errors.Is(err, domain.BadRequest):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zipcode"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting zipcode"})
		}
		return
	}

	c.JSON(200, address)

}
