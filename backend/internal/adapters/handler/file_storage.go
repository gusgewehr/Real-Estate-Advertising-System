package handler

import (
	"fmt"
	"net/http"
	"real-state-api/internal/application/port"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileStorageHandler struct {
	port             port.FileStorageInputPort
	logger           *zap.SugaredLogger
	maxUploadSize    int64
	allowedMIMETypes []string
}

func NewFileStorageHandler(port port.FileStorageInputPort, logger *zap.SugaredLogger, maxUploadSize int64, allowedMIMETypes []string) *FileStorageHandler {
	return &FileStorageHandler{port: port, logger: logger, maxUploadSize: maxUploadSize * 1024 * 1024, allowedMIMETypes: allowedMIMETypes}
}

// Post
// @Summary Upload an image
// @Schemes
// @Description Upload an image
// @Tags File Storage
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"\
// @Success 200 {object} string "Success"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Failure 400 {object} domain.ErrorResponse "Bad Request"
// @Failure 404 {object} domain.ErrorResponse "Endpoint not found"
// @Router /real-estate/image [post]
func (h *FileStorageHandler) Upload(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.maxUploadSize)

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field 'image' is required"})
		return
	}

	var fileExtension string
	contentType := fileHeader.Header.Get("Content-Type")
	for _, mimetype := range h.allowedMIMETypes {
		if strings.Contains(contentType, mimetype) {
			fileExtension = mimetype
			break
		}
	}
	if fileExtension == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("type not allowed: %s (allowed: JPEG, PNG, WebP)", contentType)})
		return
	}

	if fileHeader.Size > h.maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("exceeded upload size max: %d", h.maxUploadSize)})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		h.logger.Errorw("Failed to open file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error processing image"})
		return
	}
	defer file.Close()

	imageUrl, err := h.port.Upload(file, fileHeader.Filename)
	if err != nil {
		h.logger.Errorw("Failed to upload image", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, imageUrl)

}
