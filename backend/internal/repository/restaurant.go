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
		`INSERT INTO restaurants (user_id, name, name_en, slug, description, description_en, phone, address, working_hours, theme, cover_image, promo_title, promo_description, social_links)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, COALESCE(NULLIF($10,''), 'default'), $11, $12, $13, COALESCE(NULLIF($14,'')::jsonb, '[]'))
		 RETURNING *`,
		userID, req.Name, nullStr(req.NameEN), req.Slug,
		nullStr(req.Description), nullStr(req.DescriptionEN),
		nullStr(req.Phone), nullStr(req.Address), nullStr(req.WorkingHours), req.Theme,
		nullStr(req.CoverImage), nullStr(req.PromoTitle), nullStr(req.PromoDescription),
		nullStr(req.SocialLinks),
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

func (r *RestaurantRepo) ListAll() ([]model.Restaurant, error) {
	var list []model.Restaurant
	err := r.db.Select(&list, `SELECT * FROM restaurants ORDER BY created_at DESC`)
	return list, err
}

func (r *RestaurantRepo) Update(id int64, req *model.UpdateRestaurantRequest) (*model.Restaurant, error) {
	rest := &model.Restaurant{}
	err := r.db.QueryRowx(
		`UPDATE restaurants SET
			name = COALESCE($2, name),
			name_en = COALESCE($3, name_en),
			slug = COALESCE($4, slug),
			description = COALESCE($5, description),
			description_en = COALESCE($6, description_en),
			logo = COALESCE($7, logo),
			phone = COALESCE($8, phone),
			address = COALESCE($9, address),
			working_hours = COALESCE($10, working_hours),
			theme = COALESCE($11, theme),
			cover_image = COALESCE($12, cover_image),
			promo_title = COALESCE($13, promo_title),
			promo_description = COALESCE($14, promo_description),
			social_links = COALESCE($15::jsonb, social_links)
		 WHERE id = $1 RETURNING *`,
		id,
		emptyPtrToNil(req.Name), req.NameEN,
		emptyPtrToNil(req.Slug),
		emptyPtrToNil(req.Description), req.DescriptionEN,
		emptyPtrToNil(req.Logo),
		emptyPtrToNil(req.Phone), emptyPtrToNil(req.Address),
		emptyPtrToNil(req.WorkingHours), emptyPtrToNil(req.Theme),
		emptyPtrToNil(req.CoverImage), emptyPtrToNil(req.PromoTitle),
		emptyPtrToNil(req.PromoDescription),
		emptyPtrToNil(req.SocialLinks),
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
