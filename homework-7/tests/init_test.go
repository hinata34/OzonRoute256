//go:build integration
// +build integration

package tests

import (
	"homework-7/tests/http_server"
	"homework-7/tests/postgres"
)

var (
	Db         *postgres.TDB
	UserServer *http_server.UserServer
)

func init() {
	Db = postgres.NewTDB()
	UserServer = http_server.NewUserServer(Db)
}
