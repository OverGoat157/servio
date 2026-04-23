package repository

import (
	"github.com/jmoiron/sqlx"
	"ab-team/internal/model"
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
		`INSERT INTO menu_items (category_id, name, name_en, description, description_en, price, image, available, sort_order, ingredients, ingredients_en, weight, calories, proteins, fats, carbs, cook_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING *`,
		categoryID,
		req.Name, nullStr(req.NameEN),
		nullStr(req.Description), nullStr(req.DescriptionEN),
		req.Price,
		nullStr(req.Image), available, req.SortOrder,
		nullStr(req.Ingredients), nullStr(req.IngredientsEN),
		nullStr(req.Weight), req.Calories, req.Proteins, req.Fats, req.Carbs, nullStr(req.CookTime),
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
			name_en = COALESCE($3, name_en),
			description = COALESCE($4, description),
			description_en = COALESCE($5, description_en),
			price = COALESCE($6, price),
			image = COALESCE($7, image),
			available = COALESCE($8, available),
			sort_order = COALESCE($9, sort_order),
			ingredients = COALESCE($10, ingredients),
			ingredients_en = COALESCE($11, ingredients_en),
			weight = COALESCE($12, weight),
			calories = COALESCE($13, calories),
			proteins = COALESCE($14, proteins),
			fats = COALESCE($15, fats),
			carbs = COALESCE($16, carbs),
			cook_time = COALESCE($17, cook_time)
		 WHERE id = $1 RETURNING *`,
		id,
		req.Name, req.NameEN,
		req.Description, req.DescriptionEN,
		req.Price, req.Image, req.Available, req.SortOrder,
		req.Ingredients, req.IngredientsEN,
		req.Weight, req.Calories, req.Proteins, req.Fats, req.Carbs, req.CookTime,
	).StructScan(item)
	return item, err
}

func (r *MenuItemRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM menu_items WHERE id = $1`, id)
	return err
}
