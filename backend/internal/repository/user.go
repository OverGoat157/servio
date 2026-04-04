package repository

import (
	"github.com/jmoiron/sqlx"
	"servio/internal/model"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(email, passwordHash, name string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowx(
		`INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3)
		 RETURNING id, email, password_hash, name, created_at`,
		email, passwordHash, name,
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
