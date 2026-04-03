package model

type MessengerConfig struct {
	ID           int64  `json:"id" db:"id"`
	RestaurantID int64  `json:"restaurant_id" db:"restaurant_id"`
	Type         string `json:"type" db:"type"`       // telegram, whatsapp, viber
	Config       string `json:"config" db:"config"`   // JSON: {bot_token, chat_id} или {phone}
	Enabled      bool   `json:"enabled" db:"enabled"`
}

type CreateMessengerConfigRequest struct {
	Type    string `json:"type" binding:"required"`
	Config  string `json:"config" binding:"required"`
	Enabled *bool  `json:"enabled"`
}
