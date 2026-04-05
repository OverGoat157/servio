package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ab-team/internal/middleware"
	"ab-team/internal/model"
	"ab-team/internal/repository"
)

type RestaurantHandler struct {
	restaurants *repository.RestaurantRepo
}

func NewRestaurantHandler(restaurants *repository.RestaurantRepo) *RestaurantHandler {
	return &RestaurantHandler{restaurants: restaurants}
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var req model.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rest, err := h.restaurants.Create(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create restaurant: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rest)
}

func (h *RestaurantHandler) List(c *gin.Context) {
	userID := middleware.GetUserID(c)
	list, err := h.restaurants.ListByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list restaurants"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *RestaurantHandler) Get(c *gin.Context) {
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

	userID := middleware.GetUserID(c)
	if rest.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	c.JSON(http.StatusOK, rest)
}

func (h *RestaurantHandler) Update(c *gin.Context) {
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

	userID := middleware.GetUserID(c)
	if rest.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	var req model.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.restaurants.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *RestaurantHandler) Delete(c *gin.Context) {
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

	userID := middleware.GetUserID(c)
	if rest.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	if err := h.restaurants.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
