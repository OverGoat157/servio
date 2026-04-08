package model

type MenuItem struct {
	ID          int64    `json:"id" db:"id"`
	CategoryID  int64    `json:"category_id" db:"category_id"`
	Name        string   `json:"name" db:"name"`
	Description *string  `json:"description" db:"description"`
	Price       int      `json:"price" db:"price"` // в копейках
	Image       *string  `json:"image" db:"image"`
	Available   bool     `json:"available" db:"available"`
	SortOrder   int      `json:"sort_order" db:"sort_order"`
	Ingredients *string  `json:"ingredients" db:"ingredients"`
	Weight      *string  `json:"weight" db:"weight"`
	Calories    *int     `json:"calories" db:"calories"`
	Proteins    *float64 `json:"proteins" db:"proteins"`
	Fats        *float64 `json:"fats" db:"fats"`
	Carbs       *float64 `json:"carbs" db:"carbs"`
	CookTime    *string  `json:"cook_time" db:"cook_time"`
}

type CreateMenuItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       int     `json:"price" binding:"required,min=1"`
	Image       string  `json:"image"`
	Available   *bool   `json:"available"`
	SortOrder   int     `json:"sort_order"`
	Ingredients string  `json:"ingredients"`
	Weight      string  `json:"weight"`
	Calories    *int    `json:"calories"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
	Carbs       *float64 `json:"carbs"`
	CookTime    string  `json:"cook_time"`
}

type UpdateMenuItemRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *int     `json:"price"`
	Image       *string  `json:"image"`
	Available   *bool    `json:"available"`
	SortOrder   *int     `json:"sort_order"`
	Ingredients *string  `json:"ingredients"`
	Weight      *string  `json:"weight"`
	Calories    *int     `json:"calories"`
	Proteins    *float64 `json:"proteins"`
	Fats        *float64 `json:"fats"`
	Carbs       *float64 `json:"carbs"`
	CookTime    *string  `json:"cook_time"`
}
