package model

import "time"

type Restaurant struct {
	ID               int64     `json:"id" db:"id"`
	UserID           int64     `json:"user_id" db:"user_id"`
	Name             string    `json:"name" db:"name"`
	NameEN           *string   `json:"name_en" db:"name_en"`
	Slug             string    `json:"slug" db:"slug"`
	Description      *string   `json:"description" db:"description"`
	DescriptionEN    *string   `json:"description_en" db:"description_en"`
	Logo             *string   `json:"logo" db:"logo"`
	CoverImage       *string   `json:"cover_image" db:"cover_image"`
	Phone            *string   `json:"phone" db:"phone"`
	Address          *string   `json:"address" db:"address"`
	WorkingHours     *string   `json:"working_hours" db:"working_hours"`
	Theme            string    `json:"theme" db:"theme"`
	PromoTitle       *string   `json:"promo_title" db:"promo_title"`
	PromoDescription *string   `json:"promo_description" db:"promo_description"`
	SocialLinks      *string   `json:"social_links" db:"social_links"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

type CreateRestaurantRequest struct {
	Name             string `json:"name" binding:"required"`
	NameEN           string `json:"name_en"`
	Slug             string `json:"slug" binding:"required"`
	Description      string `json:"description"`
	DescriptionEN    string `json:"description_en"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	WorkingHours     string `json:"working_hours"`
	Theme            string `json:"theme"`
	CoverImage       string `json:"cover_image"`
	PromoTitle       string `json:"promo_title"`
	PromoDescription string `json:"promo_description"`
	SocialLinks      string `json:"social_links"`
}

type UpdateRestaurantRequest struct {
	Name             *string `json:"name"`
	NameEN           *string `json:"name_en"`
	Slug             *string `json:"slug"`
	Description      *string `json:"description"`
	DescriptionEN    *string `json:"description_en"`
	Logo             *string `json:"logo"`
	Phone            *string `json:"phone"`
	Address          *string `json:"address"`
	WorkingHours     *string `json:"working_hours"`
	Theme            *string `json:"theme"`
	CoverImage       *string `json:"cover_image"`
	PromoTitle       *string `json:"promo_title"`
	PromoDescription *string `json:"promo_description"`
	SocialLinks      *string `json:"social_links"`
}
