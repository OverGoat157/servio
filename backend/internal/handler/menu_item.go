package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"servio/internal/middleware"
	"servio/internal/model"
	"servio/internal/repository"
)

type MenuItemHandler struct {
	items       *repository.MenuItemRepo
	categories  *repository.CategoryRepo
	restaurants *repository.RestaurantRepo
}

func NewMenuItemHandler(items *repository.MenuItemRepo, categories *repository.CategoryRepo, restaurants *repository.RestaurantRepo) *MenuItemHandler {
	return &MenuItemHandler{items: items, categories: categories, restaurants: restaurants}
}

func (h *MenuItemHandler) checkCategoryOwnership(c *gin.Context, categoryID int64) bool {
	cat, err := h.categories.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return false
	}
	rest, err := h.restaurants.GetByID(cat.RestaurantID)
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

func (h *MenuItemHandler) Create(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("categoryId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
		return
	}

	if !h.checkCategoryOwnership(c, categoryID) {
		return
	}

	var req model.CreateMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.items.Create(categoryID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *MenuItemHandler) List(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("categoryId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
		return
	}

	if !h.checkCategoryOwnership(c, categoryID) {
		return
	}

	list, err := h.items.ListByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list items"})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *MenuItemHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	item, err := h.items.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	if !h.checkCategoryOwnership(c, item.CategoryID) {
		return
	}

	var req model.UpdateMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.items.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *MenuItemHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	item, err := h.items.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	if !h.checkCategoryOwnership(c, item.CategoryID) {
		return
	}

	if err := h.items.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
