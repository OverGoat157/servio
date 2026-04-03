package repository

import (
	"github.com/jmoiron/sqlx"
	"menulink/internal/model"
)

type MessengerRepo struct {
	db *sqlx.DB
}

func NewMessengerRepo(db *sqlx.DB) *MessengerRepo {
	return &MessengerRepo{db: db}
}

func (r *MessengerRepo) Upsert(restaurantID int64, req *model.CreateMessengerConfigRequest) (*model.MessengerConfig, error) {
	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}
	mc := &model.MessengerConfig{}
	err := r.db.QueryRowx(
		`INSERT INTO messenger_configs (restaurant_id, type, config, enabled)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (restaurant_id, type) DO UPDATE SET config = $3, enabled = $4
		 RETURNING *`,
		restaurantID, req.Type, req.Config, enabled,
	).StructScan(mc)
	return mc, err
}

func (r *MessengerRepo) ListByRestaurant(restaurantID int64) ([]model.MessengerConfig, error) {
	var list []model.MessengerConfig
	err := r.db.Select(&list,
		`SELECT * FROM messenger_configs WHERE restaurant_id = $1`, restaurantID)
	return list, err
}

func (r *MessengerRepo) GetEnabled(restaurantID int64, messengerType string) (*model.MessengerConfig, error) {
	mc := &model.MessengerConfig{}
	err := r.db.Get(mc,
		`SELECT * FROM messenger_configs WHERE restaurant_id = $1 AND type = $2 AND enabled = true`,
		restaurantID, messengerType)
	return mc, err
}

func (r *MessengerRepo) Delete(restaurantID int64, messengerType string) error {
	_, err := r.db.Exec(
		`DELETE FROM messenger_configs WHERE restaurant_id = $1 AND type = $2`,
		restaurantID, messengerType)
	return err
}
