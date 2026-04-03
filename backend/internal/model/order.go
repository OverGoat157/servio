package model

import "time"

type OrderItem struct {
	MenuItemID int64  `json:"menu_item_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
}

type Order struct {
	ID           int64     `json:"id" db:"id"`
	RestaurantID int64     `json:"restaurant_id" db:"restaurant_id"`
	Items        string    `json:"items" db:"items"` // JSON
	Total        int       `json:"total" db:"total"` // в копейках
	Messenger    string    `json:"messenger" db:"messenger"`
	CustomerName string    `json:"customer_name" db:"customer_name"`
	CustomerPhone string   `json:"customer_phone" db:"customer_phone"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateOrderRequest struct {
	Items         []OrderItem `json:"items" binding:"required,min=1"`
	Messenger     string      `json:"messenger" binding:"required"`
	CustomerName  string      `json:"customer_name"`
	CustomerPhone string      `json:"customer_phone"`
}
