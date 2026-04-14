package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"ab-team/internal/model"
	"ab-team/internal/repository"
)

type AdminHandler struct {
	users       *repository.UserRepo
	restaurants *repository.RestaurantRepo
	analytics   *repository.AnalyticsRepo
}

func NewAdminHandler(users *repository.UserRepo, restaurants *repository.RestaurantRepo, analytics *repository.AnalyticsRepo) *AdminHandler {
	return &AdminHandler{users: users, restaurants: restaurants, analytics: analytics}
}

// ListUsers returns all users
func (h *AdminHandler) ListUsers(c *gin.Context) {
	list, err := h.users.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetUser returns a single user by ID
func (h *AdminHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.users.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user (admin only)
func (h *AdminHandler) CreateUser(c *gin.Context) {
	var req model.AdminCreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	role := req.Role
	if role == "" {
		role = "user"
	}
	tier := req.Tier
	if tier == "" {
		tier = model.TierBasic
	} else if !model.IsValidTier(tier) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tier"})
		return
	}
	maxRest := 3
	if req.MaxRestaurants != nil {
		maxRest = *req.MaxRestaurants
	}

	user, err := h.users.Create(req.Email, string(hash), req.Name, role, tier, maxRest)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates user fields
func (h *AdminHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if _, err := h.users.GetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var req model.AdminUpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password if provided
	var passwordHash *string
	if req.Password != nil && *req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		h := string(hash)
		passwordHash = &h
	}

	if req.Tier != nil && !model.IsValidTier(*req.Tier) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tier"})
		return
	}

	user, err := h.users.Update(id, req.Email, passwordHash, req.Name, req.Role, req.Tier, req.MaxRestaurants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user and all their data
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.users.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// PlatformAnalytics returns platform-wide analytics for admin
func (h *AdminHandler) PlatformAnalytics(c *gin.Context) {
	days := 30
	if d := c.Query("days"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil && parsed > 0 && parsed <= 365 {
			days = parsed
		}
	}

	summary, err := h.analytics.GetPlatformSummary(days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load analytics"})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// ListAllRestaurants returns all restaurants across all users
func (h *AdminHandler) ListAllRestaurants(c *gin.Context) {
	list, err := h.restaurants.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list restaurants"})
		return
	}
	c.JSON(http.StatusOK, list)
}

// UpdateAnyRestaurant updates any restaurant regardless of owner
func (h *AdminHandler) UpdateAnyRestaurant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if _, err := h.restaurants.GetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
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
