package address

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type AddressRepo interface {
	Add(ctx context.Context, address *Address) (int64, error)
	GetById(ctx context.Context, number int64) (*Address, error)
	List(ctx context.Context) ([]*Address, error)
	Update(ctx context.Context, address *Address) (bool, error)
}
