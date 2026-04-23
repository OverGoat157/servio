package repository

import (
	"github.com/jmoiron/sqlx"

	"ab-team/internal/model"
)

type ComboRepo struct {
	db *sqlx.DB
}

func NewComboRepo(db *sqlx.DB) *ComboRepo {
	return &ComboRepo{db: db}
}

func (r *ComboRepo) Create(restaurantID int64, req *model.CreateComboRequest) (*model.Combo, error) {
	available := true
	if req.Available != nil {
		available = *req.Available
	}
	combo := &model.Combo{}
	err := r.db.QueryRowx(
		`INSERT INTO combos (restaurant_id, name, name_en, description, description_en, image, price, available, sort_order)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`,
		restaurantID,
		req.Name, nullStr(req.NameEN),
		nullStr(req.Description), nullStr(req.DescriptionEN),
		nullStr(req.Image),
		req.Price, available, req.SortOrder,
	).StructScan(combo)
	if err != nil {
		return nil, err
	}

	for _, ci := range req.Items {
		qty := ci.Quantity
		if qty < 1 {
			qty = 1
		}
		_, err := r.db.Exec(
			`INSERT INTO combo_items (combo_id, menu_item_id, name, name_en, quantity) VALUES ($1, $2, $3, $4, $5)`,
			combo.ID, ci.MenuItemID, ci.Name, nullStr(ci.NameEN), qty,
		)
		if err != nil {
			return nil, err
		}
	}

	return combo, nil
}

func (r *ComboRepo) ListByRestaurant(restaurantID int64) ([]model.Combo, error) {
	var list []model.Combo
	err := r.db.Select(&list,
		`SELECT * FROM combos WHERE restaurant_id = $1 ORDER BY sort_order, id`, restaurantID)
	return list, err
}

func (r *ComboRepo) GetByID(id int64) (*model.Combo, error) {
	combo := &model.Combo{}
	err := r.db.Get(combo, `SELECT * FROM combos WHERE id = $1`, id)
	return combo, err
}

func (r *ComboRepo) GetItems(comboID int64) ([]model.ComboItem, error) {
	var items []model.ComboItem
	err := r.db.Select(&items, `SELECT * FROM combo_items WHERE combo_id = $1 ORDER BY id`, comboID)
	return items, err
}

func (r *ComboRepo) Update(id int64, req *model.UpdateComboRequest) (*model.Combo, error) {
	combo := &model.Combo{}
	err := r.db.QueryRowx(
		`UPDATE combos SET
			name = COALESCE($2, name),
			name_en = COALESCE($3, name_en),
			description = COALESCE($4, description),
			description_en = COALESCE($5, description_en),
			image = COALESCE($6, image),
			price = COALESCE($7, price),
			available = COALESCE($8, available),
			sort_order = COALESCE($9, sort_order)
		 WHERE id = $1 RETURNING *`,
		id,
		req.Name, req.NameEN,
		req.Description, req.DescriptionEN,
		req.Image, req.Price, req.Available, req.SortOrder,
	).StructScan(combo)
	if err != nil {
		return nil, err
	}

	if req.Items != nil {
		// Replace all items
		r.db.Exec(`DELETE FROM combo_items WHERE combo_id = $1`, id)
		for _, ci := range req.Items {
			qty := ci.Quantity
			if qty < 1 {
				qty = 1
			}
			r.db.Exec(
				`INSERT INTO combo_items (combo_id, menu_item_id, name, name_en, quantity) VALUES ($1, $2, $3, $4, $5)`,
				id, ci.MenuItemID, ci.Name, nullStr(ci.NameEN), qty,
			)
		}
	}

	return combo, nil
}

func (r *ComboRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM combos WHERE id = $1`, id)
	return err
}
