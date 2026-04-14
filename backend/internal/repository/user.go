package repository

import (
	"github.com/jmoiron/sqlx"
	"ab-team/internal/model"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(email, passwordHash, name, role, tier string, maxRestaurants int) (*model.User, error) {
	if role == "" {
		role = "user"
	}
	if tier == "" {
		tier = model.TierBasic
	}
	if maxRestaurants == 0 {
		maxRestaurants = 3
	}
	user := &model.User{}
	err := r.db.QueryRowx(
		`INSERT INTO users (email, password_hash, name, role, tier, max_restaurants)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		email, passwordHash, name, role, tier, maxRestaurants,
	).StructScan(user)
	return user, err
}

func (r *UserRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Get(user, `SELECT * FROM users WHERE email = $1`, email)
	return user, err
}

func (r *UserRepo) GetByID(id int64) (*model.User, error) {
	user := &model.User{}
	err := r.db.Get(user, `SELECT * FROM users WHERE id = $1`, id)
	return user, err
}

func (r *UserRepo) ListAll() ([]model.User, error) {
	var list []model.User
	err := r.db.Select(&list, `SELECT * FROM users ORDER BY created_at DESC`)
	return list, err
}

func (r *UserRepo) Update(id int64, email, passwordHash, name, role, tier *string, maxRestaurants *int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowx(
		`UPDATE users SET
			email = COALESCE($2, email),
			password_hash = COALESCE($3, password_hash),
			name = COALESCE($4, name),
			role = COALESCE($5, role),
			tier = COALESCE($6, tier),
			max_restaurants = COALESCE($7, max_restaurants)
		 WHERE id = $1 RETURNING *`,
		id, email, passwordHash, name, role, tier, maxRestaurants,
	).StructScan(user)
	return user, err
}

func (r *UserRepo) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func (r *UserRepo) CountRestaurants(userID int64) (int, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM restaurants WHERE user_id = $1`, userID)
	return count, err
}
