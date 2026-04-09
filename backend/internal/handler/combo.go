package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ab-team/internal/middleware"
	"ab-team/internal/model"
	"ab-team/internal/repository"
)

type ComboHandler struct {
	combos      *repository.ComboRepo
	restaurants *repository.RestaurantRepo
}

func NewComboHandler(combos *repository.ComboRepo, restaurants *repository.RestaurantRepo) *ComboHandler {
	return &ComboHandler{combos: combos, restaurants: restaurants}
}

func (h *ComboHandler) checkOwnership(c *gin.Context, restaurantID int64) bool {
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

func (h *ComboHandler) Create(c *gin.Context) {
	restID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}
	if !h.checkOwnership(c, restID) {
		return
	}

	var req model.CreateComboRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combo, err := h.combos.Create(restID, &req)
	if err != nil {
		if isDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Комбо с таким названием уже существует"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create combo"})
		return
	}

	c.JSON(http.StatusCreated, combo)
}

type ComboWithItems struct {
	model.Combo
	Items []model.ComboItem `json:"items"`
}

func (h *ComboHandler) List(c *gin.Context) {
	restID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}
	if !h.checkOwnership(c, restID) {
		return
	}

	combos, err := h.combos.ListByRestaurant(restID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list combos"})
		return
	}

	var result []ComboWithItems
	for _, combo := range combos {
		items, _ := h.combos.GetItems(combo.ID)
		if items == nil {
			items = []model.ComboItem{}
		}
		result = append(result, ComboWithItems{Combo: combo, Items: items})
	}

	if result == nil {
		result = []ComboWithItems{}
	}
	c.JSON(http.StatusOK, result)
}

func (h *ComboHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	combo, err := h.combos.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "combo not found"})
		return
	}

	if !h.checkOwnership(c, combo.RestaurantID) {
		return
	}

	var req model.UpdateComboRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Image != nil && combo.Image != nil && *req.Image != *combo.Image {
		DeleteFile(*combo.Image)
	}

	updated, err := h.combos.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update combo"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *ComboHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	combo, err := h.combos.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "combo not found"})
		return
	}

	if !h.checkOwnership(c, combo.RestaurantID) {
		return
	}

	if combo.Image != nil {
		DeleteFile(*combo.Image)
	}

	if err := h.combos.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete combo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
