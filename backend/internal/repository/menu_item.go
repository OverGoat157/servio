package repository

import (
	"github.com/jmoiron/sqlx"
	"menulink/internal/model"
)

type MenuItemRepo struct {
	db *sqlx.DB
}

func NewMenuItemRepo(db *sqlx.DB) *MenuItemRepo {
	return &MenuItemRepo{db: db}
}

func (r *MenuItemRepo) Create(categoryID int64, req *model.CreateMenuItemRequest) (*model.MenuItem, error) {
	available := true
	if req.Available != nil {
		available = *req.Available
	}
	item := &model.MenuItem{}
	err := r.db.QueryRowx(
		`INSERT INTO menu_items (category_id, name, description, price, image, available, sort_order)
		 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		categoryID, req.Name, nullStr(req.Description), req.Price,
		nullStr(req.Image), available, req.SortOrder,
	).StructScan(item)
	return item, err
}

func (r *MenuItemRepo) ListByCategory(categoryID int64) ([]model.MenuItem, error) {
	var list []model.MenuItem
	err := r.db.Select(&list,
		`SELECT * FROM menu_items WHERE category_id = $1 ORDER BY sort_order, id`, categoryID)
	return list, err
}

func (r *MenuItemRepo) ListByRestaurant(restaurantID int64) ([]model.MenuItem, error) {
	var list []model.MenuItem
	err := r.db.Select(&list,
		`SELECT mi.* FROM menu_items mi
		 JOIN categories c ON c.id = mi.category_id
		 WHERE c.restaurant_id = $1
		 ORDER BY c.sort_order, c.id, mi.sort_order, mi.id`, restaurantID)
	return list, err
}

func (r *MenuItemRepo) GetByID(id int64) (*model.MenuItem, error) {
	item := &model.MenuItem{}
	err := r.db.Get(item, `SELECT * FROM menu_items WHERE id = $1`, id)
	return item, err
}

func (r *MenuItemRepo) Update(id int64, req *model.UpdateMenuItemRequest) (*model.MenuItem, error) {
	item := &model.MenuItem{}
	err := r.db.QueryRowx(
		`UPDATE menu_items SET
			name = COALESCE($2, name),
			description = COALESCE($3, description),
			price = COALESCE($4, price),
			image = COALESCE($5, image),
			available = COALESCE($6, available),
			sort_order = COALESCE($7, sort_order)
		 WHERE id = $1 RETURNING *`,
		id, req.Name, req.Description, req.Price, req.Image, req.Available, req.SortOrder,
	).StructScan(item)
	return item, err
}

func (r *MenuItemRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM menu_items WHERE id = $1`, id)
	return err
}
