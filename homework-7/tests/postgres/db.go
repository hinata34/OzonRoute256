//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"homework-7/internal/app/db"
	"strings"
	"sync"
	"testing"
)

type TDB struct {
	sync.Mutex
	DB *db.Database
}

func NewTDB() *TDB {
	ctx := context.Background()
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "test", "test", "test")
	db, _ := db.NewDatabase(ctx, psqlConn)
	return &TDB{DB: db}
}

func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)
}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (d *TDB) Truncate(ctx context.Context) {
	tables := []string{"users", "addresses"}
	q := fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ","))
	_, err := d.DB.Exec(ctx, q)
	if err != nil {
		panic(err)
	}
}
