package main

import (
	"context"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/processinput"
	"homework-5/internal/pkg/repository/postgresql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		return
	}
	defer database.GetPool(ctx).Close()

	usersRepo := postgresql.NewUsers(database)
	addressesRepo := postgresql.NewAddresses(database)

	processinput.DBInput(ctx, usersRepo, addressesRepo)
}
