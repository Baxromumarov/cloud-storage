package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) UploadFile(c *gin.Context) {
	start := time.Now()

	file, err := c.FormFile("file")
	if err != nil {
		h.log.Error("Error while getting file from request", logger.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	fileInfo := fmt.Sprintf("File original name: %s, Size: %f MB", file.Filename, float64(file.Size)/1024/1024)
	h.log.Info(fileInfo)

	dst, err := os.Create(file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Create file err: %s", err.Error()))
		return
	}
	defer dst.Close()

	nfile, err := file.Open()
	if err != nil {
		h.log.Error("Error while opening file", logger.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Open file err: %s", err.Error()))
		return
	}

	if _, err := io.Copy(dst, nfile); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file err: %s", err.Error()))
		return
	}
	fmt.Println(time.Since(start))

	h.log.Info("File uploaded successfully", logger.Duration("duration", time.Since(start)))

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
	})
	// return

}
