package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ab-team/internal/middleware"
	"ab-team/internal/repository"
)

type OrderHandler struct {
	orders      *repository.OrderRepo
	restaurants *repository.RestaurantRepo
}

func NewOrderHandler(orders *repository.OrderRepo, restaurants *repository.RestaurantRepo) *OrderHandler {
	return &OrderHandler{orders: orders, restaurants: restaurants}
}

func (h *OrderHandler) List(c *gin.Context) {
	restaurantID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant_id"})
		return
	}

	rest, err := h.restaurants.GetByID(restaurantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	if rest.UserID != middleware.GetUserID(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	limit := 50
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil && v >= 0 {
			offset = v
		}
	}

	orders, err := h.orders.ListByRestaurant(restaurantID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	order, err := h.orders.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	rest, err := h.restaurants.GetByID(order.RestaurantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	if rest.UserID != middleware.GetUserID(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orders.UpdateStatus(id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}
