package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	uploadsDir  = "uploads"
	maxFileSize = 10 << 20 // 10 MB
)

var allowedExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".gif": true, ".svg": true,
}

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	os.MkdirAll(uploadsDir, 0755)
	return &UploadHandler{}
}

// Upload handles POST /api/upload with multipart file
func (h *UploadHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "файл не найден"})
		return
	}

	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "файл слишком большой (макс. 10 МБ)"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый формат файла"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join(uploadsDir, filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить файл"})
		return
	}

	url := fmt.Sprintf("/uploads/%s", filename)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// DeleteFile removes an old uploaded file by its URL path
func DeleteFile(urlPath string) {
	if urlPath == "" || !strings.HasPrefix(urlPath, "/uploads/") {
		return
	}
	filename := filepath.Base(urlPath)
	path := filepath.Join(uploadsDir, filename)
	os.Remove(path)
}
