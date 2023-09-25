package address

import (
	"context"
	"homework-5/internal/app/db"
)

type AddressesRepo struct {
	db db.DBops
}

func NewAddresses(db db.DBops) *AddressesRepo {
	return &AddressesRepo{db: db}
}

func (r *AddressesRepo) Add(ctx context.Context, address *Address) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO addresses(house_number, street_name, user_id) VALUES ($1, $2, $3) RETURNING ID", address.HouseNumber, address.StreetName, address.UserID).Scan(&id)
	return id, err
}

func (r *AddressesRepo) GetById(ctx context.Context, id int64) (*Address, error) {
	var a Address
	err := r.db.Get(ctx, &a, "SELECT id, house_number, street_name, user_id FROM addresses WHERE id = $1", id)
	return &a, err
}

func (r *AddressesRepo) List(ctx context.Context) ([]*Address, error) {
	addresses := make([]*Address, 0)
	err := r.db.Select(ctx, &addresses, "SELECT id, house_number, street_name, user_id FROM addresses")
	return addresses, err
}

func (r *AddressesRepo) Update(ctx context.Context, address *Address) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE addresses SET house_number = $1, street_name = $2, user_id = $3 WHERE id = $4", address.HouseNumber, address.StreetName, address.UserID, address.ID)
	return result.RowsAffected() > 0, err
}
