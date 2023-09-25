package postgresql

import (
	"context"
	"database/sql"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
)

type UsersRepo struct {
	db db.DBops
}

func NewUsers(db db.DBops) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO users(name, age) VALUES ($1, $2) RETURNING id", user.Name, user.Age).Scan(&id)
	return id, err
}

func (r *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	var u repository.User
	err := r.db.Get(ctx, &u, "SELECT id,name,age FROM users WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	return &u, err
}

func (r *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	users := make([]*repository.User, 0)
	err := r.db.Select(ctx, &users, "SELECT id,name,age FROM users")
	return users, err
}

func (r *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE users SET name = $1, age = $2 WHERE id = $3", user.Name, user.Age, user.ID)
	return result.RowsAffected() > 0, err
}
