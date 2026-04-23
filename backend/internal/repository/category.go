package repository

import (
	"github.com/jmoiron/sqlx"
	"ab-team/internal/model"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(restaurantID int64, req *model.CreateCategoryRequest) (*model.Category, error) {
	cat := &model.Category{}
	err := r.db.QueryRowx(
		`INSERT INTO categories (restaurant_id, name, name_en, sort_order) VALUES ($1, $2, $3, $4) RETURNING *`,
		restaurantID, req.Name, nullStr(req.NameEN), req.SortOrder,
	).StructScan(cat)
	return cat, err
}

func (r *CategoryRepo) ListByRestaurant(restaurantID int64) ([]model.Category, error) {
	var list []model.Category
	err := r.db.Select(&list,
		`SELECT * FROM categories WHERE restaurant_id = $1 ORDER BY sort_order, id`, restaurantID)
	return list, err
}

func (r *CategoryRepo) GetByID(id int64) (*model.Category, error) {
	cat := &model.Category{}
	err := r.db.Get(cat, `SELECT * FROM categories WHERE id = $1`, id)
	return cat, err
}

func (r *CategoryRepo) Update(id int64, req *model.UpdateCategoryRequest) (*model.Category, error) {
	cat := &model.Category{}
	err := r.db.QueryRowx(
		`UPDATE categories SET
			name = COALESCE($2, name),
			name_en = COALESCE($3, name_en),
			sort_order = COALESCE($4, sort_order)
		 WHERE id = $1 RETURNING *`,
		id, req.Name, req.NameEN, req.SortOrder,
	).StructScan(cat)
	return cat, err
}

func (r *CategoryRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM categories WHERE id = $1`, id)
	return err
}
