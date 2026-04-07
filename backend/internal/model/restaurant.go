package model

import "time"

type Restaurant struct {
	ID               int64     `json:"id" db:"id"`
	UserID           int64     `json:"user_id" db:"user_id"`
	Name             string    `json:"name" db:"name"`
	Slug             string    `json:"slug" db:"slug"`
	Description      *string   `json:"description" db:"description"`
	Logo             *string   `json:"logo" db:"logo"`
	CoverImage       *string   `json:"cover_image" db:"cover_image"`
	Phone            *string   `json:"phone" db:"phone"`
	Address          *string   `json:"address" db:"address"`
	WorkingHours     *string   `json:"working_hours" db:"working_hours"`
	Theme            string    `json:"theme" db:"theme"`
	PromoTitle       *string   `json:"promo_title" db:"promo_title"`
	PromoDescription *string   `json:"promo_description" db:"promo_description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type CreateRestaurantRequest struct {
	Name             string `json:"name" binding:"required"`
	Slug             string `json:"slug" binding:"required"`
	Description      string `json:"description"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	WorkingHours     string `json:"working_hours"`
	Theme            string `json:"theme"`
	CoverImage       string `json:"cover_image"`
	PromoTitle       string `json:"promo_title"`
	PromoDescription string `json:"promo_description"`
}

type UpdateRestaurantRequest struct {
	Name             *string `json:"name"`
	Slug             *string `json:"slug"`
	Description      *string `json:"description"`
	Logo             *string `json:"logo"`
	Phone            *string `json:"phone"`
	Address          *string `json:"address"`
	WorkingHours     *string `json:"working_hours"`
	Theme            *string `json:"theme"`
	CoverImage       *string `json:"cover_image"`
	PromoTitle       *string `json:"promo_title"`
	PromoDescription *string `json:"promo_description"`
}
