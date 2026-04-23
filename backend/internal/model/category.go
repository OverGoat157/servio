package model

type Category struct {
	ID           int64   `json:"id" db:"id"`
	RestaurantID int64   `json:"restaurant_id" db:"restaurant_id"`
	Name         string  `json:"name" db:"name"`
	NameEN       *string `json:"name_en" db:"name_en"`
	SortOrder    int     `json:"sort_order" db:"sort_order"`
}

type CreateCategoryRequest struct {
	Name      string `json:"name" binding:"required"`
	NameEN    string `json:"name_en"`
	SortOrder int    `json:"sort_order"`
}

type UpdateCategoryRequest struct {
	Name      *string `json:"name"`
	NameEN    *string `json:"name_en"`
	SortOrder *int    `json:"sort_order"`
}
