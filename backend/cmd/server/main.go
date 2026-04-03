package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"menulink/internal/config"
	"menulink/internal/handler"
	"menulink/internal/middleware"
	"menulink/internal/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := repository.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Repositories
	userRepo := repository.NewUserRepo(db)
	restRepo := repository.NewRestaurantRepo(db)
	catRepo := repository.NewCategoryRepo(db)
	itemRepo := repository.NewMenuItemRepo(db)
	orderRepo := repository.NewOrderRepo(db)
	msgRepo := repository.NewMessengerRepo(db)

	// Handlers
	authH := handler.NewAuthHandler(userRepo, cfg.JWTSecret)
	restH := handler.NewRestaurantHandler(restRepo)
	catH := handler.NewCategoryHandler(catRepo, restRepo)
	itemH := handler.NewMenuItemHandler(itemRepo, catRepo, restRepo)
	orderH := handler.NewOrderHandler(orderRepo, restRepo)
	msgH := handler.NewMessengerHandler(msgRepo, restRepo)
	publicH := handler.NewPublicHandler(restRepo, catRepo, itemRepo, orderRepo, msgRepo)

	r := gin.Default()
	r.Use(middleware.CORS())

	// --- Public API (без авторизации) ---
	pub := r.Group("/api/public")
	{
		pub.GET("/menu/:slug", publicH.GetMenu)
		pub.POST("/menu/:slug/order", publicH.CreateOrder)
	}

	// --- Auth ---
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authH.Register)
		auth.POST("/login", authH.Login)
	}

	// --- Protected API ---
	api := r.Group("/api", middleware.AuthRequired(cfg.JWTSecret))
	{
		api.GET("/me", authH.Me)

		// Restaurants
		api.POST("/restaurants", restH.Create)
		api.GET("/restaurants", restH.List)
		api.GET("/restaurants/:id", restH.Get)
		api.PUT("/restaurants/:id", restH.Update)
		api.DELETE("/restaurants/:id", restH.Delete)

		// Categories
		api.POST("/restaurants/:restaurantId/categories", catH.Create)
		api.GET("/restaurants/:restaurantId/categories", catH.List)
		api.PUT("/categories/:id", catH.Update)
		api.DELETE("/categories/:id", catH.Delete)

		// Menu items
		api.POST("/categories/:categoryId/items", itemH.Create)
		api.GET("/categories/:categoryId/items", itemH.List)
		api.PUT("/items/:id", itemH.Update)
		api.DELETE("/items/:id", itemH.Delete)

		// Orders
		api.GET("/restaurants/:restaurantId/orders", orderH.List)
		api.PATCH("/orders/:id/status", orderH.UpdateStatus)

		// Messengers
		api.GET("/restaurants/:restaurantId/messengers", msgH.List)
		api.POST("/restaurants/:restaurantId/messengers", msgH.Upsert)
		api.DELETE("/restaurants/:restaurantId/messengers/:type", msgH.Delete)
	}

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
