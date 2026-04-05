package repository

import (
	"github.com/jmoiron/sqlx"
	"ab-team/internal/model"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(restaurantID int64, itemsJSON string, total int, messenger, customerName, customerPhone string) (*model.Order, error) {
	order := &model.Order{}
	err := r.db.QueryRowx(
		`INSERT INTO orders (restaurant_id, items, total, messenger, customer_name, customer_phone)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		restaurantID, itemsJSON, total, messenger, customerName, customerPhone,
	).StructScan(order)
	return order, err
}

func (r *OrderRepo) ListByRestaurant(restaurantID int64, limit, offset int) ([]model.Order, error) {
	var list []model.Order
	err := r.db.Select(&list,
		`SELECT * FROM orders WHERE restaurant_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		restaurantID, limit, offset)
	return list, err
}

func (r *OrderRepo) GetByID(id int64) (*model.Order, error) {
	order := &model.Order{}
	err := r.db.Get(order, `SELECT * FROM orders WHERE id = $1`, id)
	return order, err
}

func (r *OrderRepo) UpdateStatus(id int64, status string) error {
	_, err := r.db.Exec(`UPDATE orders SET status = $2 WHERE id = $1`, id, status)
	return err
}
