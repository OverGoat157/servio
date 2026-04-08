package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"

	"ab-team/internal/middleware"
	"ab-team/internal/repository"
)

type QRCodeHandler struct {
	restaurants *repository.RestaurantRepo
}

func NewQRCodeHandler(restaurants *repository.RestaurantRepo) *QRCodeHandler {
	return &QRCodeHandler{restaurants: restaurants}
}

func (h *QRCodeHandler) Generate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rest, err := h.restaurants.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	if rest.UserID != middleware.GetUserID(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	menuURL := "https://menu.ab-team.ru/" + rest.Slug

	size := 512
	if s := c.Query("size"); s != "" {
		if parsed, err := strconv.Atoi(s); err == nil && parsed >= 128 && parsed <= 2048 {
			size = parsed
		}
	}

	png, err := qrcode.Encode(menuURL, qrcode.Medium, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate qr code"})
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}
