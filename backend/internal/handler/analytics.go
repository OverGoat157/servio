package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ab-team/internal/middleware"
	"ab-team/internal/repository"
)

type AnalyticsHandler struct {
	analytics   *repository.AnalyticsRepo
	restaurants *repository.RestaurantRepo
}

func NewAnalyticsHandler(analytics *repository.AnalyticsRepo, restaurants *repository.RestaurantRepo) *AnalyticsHandler {
	return &AnalyticsHandler{analytics: analytics, restaurants: restaurants}
}

func (h *AnalyticsHandler) GetSummary(c *gin.Context) {
	idStr := c.Param("id")
	restID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Проверяем что ресторан принадлежит пользователю
	userID := middleware.GetUserID(c)
	rest, err := h.restaurants.GetByID(restID)
	if err != nil || rest.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	days := 30
	if d := c.Query("days"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil && parsed > 0 && parsed <= 365 {
			days = parsed
		}
	}

	summary, err := h.analytics.GetSummary(restID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load analytics"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
