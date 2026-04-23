package model

import "time"

type Combo struct {
	ID            int64     `json:"id" db:"id"`
	RestaurantID  int64     `json:"restaurant_id" db:"restaurant_id"`
	Name          string    `json:"name" db:"name"`
	NameEN        *string   `json:"name_en" db:"name_en"`
	Description   *string   `json:"description" db:"description"`
	DescriptionEN *string   `json:"description_en" db:"description_en"`
	Image         *string   `json:"image" db:"image"`
	Price         int       `json:"price" db:"price"`
	Available     bool      `json:"available" db:"available"`
	SortOrder     int       `json:"sort_order" db:"sort_order"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type ComboItem struct {
	ID         int64   `json:"id" db:"id"`
	ComboID    int64   `json:"combo_id" db:"combo_id"`
	MenuItemID *int64  `json:"menu_item_id" db:"menu_item_id"`
	Name       string  `json:"name" db:"name"`
	NameEN     *string `json:"name_en" db:"name_en"`
	Quantity   int     `json:"quantity" db:"quantity"`
}

type CreateComboRequest struct {
	Name          string             `json:"name" binding:"required"`
	NameEN        string             `json:"name_en"`
	Description   string             `json:"description"`
	DescriptionEN string             `json:"description_en"`
	Image         string             `json:"image"`
	Price         int                `json:"price" binding:"required,min=1"`
	Available     *bool              `json:"available"`
	SortOrder     int                `json:"sort_order"`
	Items         []ComboItemRequest `json:"items"`
}

type UpdateComboRequest struct {
	Name          *string            `json:"name"`
	NameEN        *string            `json:"name_en"`
	Description   *string            `json:"description"`
	DescriptionEN *string            `json:"description_en"`
	Image         *string            `json:"image"`
	Price         *int               `json:"price"`
	Available     *bool              `json:"available"`
	SortOrder     *int               `json:"sort_order"`
	Items         []ComboItemRequest `json:"items"`
}

type ComboItemRequest struct {
	MenuItemID *int64 `json:"menu_item_id"`
	Name       string `json:"name"`
	NameEN     string `json:"name_en"`
	Quantity   int    `json:"quantity"`
}
