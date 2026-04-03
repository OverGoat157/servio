package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"menulink/internal/middleware"
	"menulink/internal/model"
	"menulink/internal/repository"
)

type MessengerHandler struct {
	messengers  *repository.MessengerRepo
	restaurants *repository.RestaurantRepo
}

func NewMessengerHandler(messengers *repository.MessengerRepo, restaurants *repository.RestaurantRepo) *MessengerHandler {
	return &MessengerHandler{messengers: messengers, restaurants: restaurants}
}

func (h *MessengerHandler) checkOwnership(c *gin.Context, restaurantID int64) bool {
	rest, err := h.restaurants.GetByID(restaurantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return false
	}
	if rest.UserID != middleware.GetUserID(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return false
	}
	return true
}

func (h *MessengerHandler) List(c *gin.Context) {
	restaurantID, err := strconv.ParseInt(c.Param("restaurantId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant_id"})
		return
	}

	if !h.checkOwnership(c, restaurantID) {
		return
	}

	list, err := h.messengers.ListByRestaurant(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list"})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *MessengerHandler) Upsert(c *gin.Context) {
	restaurantID, err := strconv.ParseInt(c.Param("restaurantId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant_id"})
		return
	}

	if !h.checkOwnership(c, restaurantID) {
		return
	}

	var req model.CreateMessengerConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mc, err := h.messengers.Upsert(restaurantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, mc)
}

func (h *MessengerHandler) Delete(c *gin.Context) {
	restaurantID, err := strconv.ParseInt(c.Param("restaurantId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant_id"})
		return
	}

	if !h.checkOwnership(c, restaurantID) {
		return
	}

	messengerType := c.Param("type")
	if err := h.messengers.Delete(restaurantID, messengerType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
