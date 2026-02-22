package handler

import (
	"net/http"
	"real-estate-api/internal/application/port"
	"real-estate-api/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RealEstateHandler struct {
	port   port.RealEstateInputPort
	logger *zap.SugaredLogger
}

func NewRealEstateHandler(port port.RealEstateInputPort, logger *zap.SugaredLogger) *RealEstateHandler {
	return &RealEstateHandler{port, logger}
}

// Post
// @Summary Create a new real estate property
// @Schemes
// @Description Create a new real estate property
// @Tags Real Estate
// @Accept json
// @Produce json
// @Success 201 {object} domain.NoBodyResponse "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"
// @Router /real-estate [post]
func (h *RealEstateHandler) CreateRealEstate(c *gin.Context) {
	var input *domain.RealEstatePropertyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Errorw("Failed to bind JSON", "erro", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	err := h.port.Create(input)
	if err != nil {
		h.logger.Errorw("Error creating real estate", "erro", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating real estate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Real estate created successfully"})
}

// Get
// @Summary List real estates
// @Schemes
// @Description List real estates
// @Tags Real Estate
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} domain.RealEstatePaginatedResponse "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"
// @Router /real-estate [get]
func (h *RealEstateHandler) ListRealEstates(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	if pageSizeNum < -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	if pageNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	response, err := h.port.List(pageNum, int64(pageSizeNum))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing real estates"})
		return
	}

	c.JSON(http.StatusOK, response)
}
