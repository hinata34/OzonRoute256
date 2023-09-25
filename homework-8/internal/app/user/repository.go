//go:generate mockgen -source=./repository.go -destination=./mocks/repository.go -package=mock_repository
package user

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type UserRepo interface {
	Add(ctx context.Context, user *User) (int64, error)
	GetById(ctx context.Context, id int64) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) (bool, error)
}
