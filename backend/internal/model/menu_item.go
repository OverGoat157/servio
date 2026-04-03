package model

type MenuItem struct {
	ID          int64   `json:"id" db:"id"`
	CategoryID  int64   `json:"category_id" db:"category_id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Price       int     `json:"price" db:"price"` // в копейках
	Image       *string `json:"image" db:"image"`
	Available   bool    `json:"available" db:"available"`
	SortOrder   int     `json:"sort_order" db:"sort_order"`
}

type CreateMenuItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required,min=1"`
	Image       string `json:"image"`
	Available   *bool  `json:"available"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateMenuItemRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	Image       *string `json:"image"`
	Available   *bool   `json:"available"`
	SortOrder   *int    `json:"sort_order"`
}
