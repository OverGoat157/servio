package repository

import (
	"github.com/jmoiron/sqlx"
	"ab-team/internal/model"
)

type RestaurantRepo struct {
	db *sqlx.DB
}

func NewRestaurantRepo(db *sqlx.DB) *RestaurantRepo {
	return &RestaurantRepo{db: db}
}

func (r *RestaurantRepo) Create(userID int64, req *model.CreateRestaurantRequest) (*model.Restaurant, error) {
	rest := &model.Restaurant{}
	err := r.db.QueryRowx(
		`INSERT INTO restaurants (user_id, name, slug, description, phone, address, working_hours, theme, cover_image, promo_title, promo_description)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, COALESCE(NULLIF($8,''), 'default'), $9, $10, $11)
		 RETURNING *`,
		userID, req.Name, req.Slug, nullStr(req.Description), nullStr(req.Phone),
		nullStr(req.Address), nullStr(req.WorkingHours), req.Theme,
		nullStr(req.CoverImage), nullStr(req.PromoTitle), nullStr(req.PromoDescription),
	).StructScan(rest)
	return rest, err
}

func (r *RestaurantRepo) GetByID(id int64) (*model.Restaurant, error) {
	rest := &model.Restaurant{}
	err := r.db.Get(rest, `SELECT * FROM restaurants WHERE id = $1`, id)
	return rest, err
}

func (r *RestaurantRepo) GetBySlug(slug string) (*model.Restaurant, error) {
	rest := &model.Restaurant{}
	err := r.db.Get(rest, `SELECT * FROM restaurants WHERE slug = $1`, slug)
	return rest, err
}

func (r *RestaurantRepo) ListByUser(userID int64) ([]model.Restaurant, error) {
	var list []model.Restaurant
	err := r.db.Select(&list, `SELECT * FROM restaurants WHERE user_id = $1 ORDER BY created_at DESC`, userID)
	return list, err
}

func (r *RestaurantRepo) Update(id int64, req *model.UpdateRestaurantRequest) (*model.Restaurant, error) {
	rest := &model.Restaurant{}
	err := r.db.QueryRowx(
		`UPDATE restaurants SET
			name = COALESCE($2, name),
			slug = COALESCE($3, slug),
			description = COALESCE($4, description),
			logo = COALESCE($5, logo),
			phone = COALESCE($6, phone),
			address = COALESCE($7, address),
			working_hours = COALESCE($8, working_hours),
			theme = COALESCE($9, theme),
			cover_image = COALESCE($10, cover_image),
			promo_title = COALESCE($11, promo_title),
			promo_description = COALESCE($12, promo_description)
		 WHERE id = $1 RETURNING *`,
		id, emptyPtrToNil(req.Name), emptyPtrToNil(req.Slug),
		emptyPtrToNil(req.Description), emptyPtrToNil(req.Logo),
		emptyPtrToNil(req.Phone), emptyPtrToNil(req.Address),
		emptyPtrToNil(req.WorkingHours), emptyPtrToNil(req.Theme),
		emptyPtrToNil(req.CoverImage), emptyPtrToNil(req.PromoTitle),
		emptyPtrToNil(req.PromoDescription),
	).StructScan(rest)
	return rest, err
}

func (r *RestaurantRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM restaurants WHERE id = $1`, id)
	return err
}

func nullStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// emptyPtrToNil converts a pointer to an empty string to nil.
// Prevents JSONB and other typed columns from receiving empty strings.
func emptyPtrToNil(s *string) *string {
	if s != nil && *s == "" {
		return nil
	}
	return s
}
