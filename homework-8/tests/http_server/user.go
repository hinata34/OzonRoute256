//go:build integration
// +build integration

package http_server

import (
	"context"
	"fmt"
	"homework-7/internal/app/server"
	"homework-7/internal/app/user"
	"homework-7/tests/postgres"
	"net/http"
	"sync"
	"testing"
)

type UserServer struct {
	sync.Mutex
	TDB  *postgres.TDB
	serv *http.Server
}

func NewUserServer(db *postgres.TDB) *UserServer {
	userRepo := user.NewUsers(db.DB)
	return &UserServer{TDB: db, serv: &http.Server{Addr: ":1337", Handler: server.CreateServer(context.Background(), userRepo)}}
}

func (u *UserServer) SetUp(t *testing.T) {
	t.Helper()
	u.Lock()
	u.TDB.SetUp(t)
	go func(u *UserServer) {
		if err := u.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
			panic("can't create server")
		}
	}(u)
}

func (u *UserServer) TearDown() {
	defer u.Unlock()
	u.serv.Shutdown(context.Background())
	u.TDB.TearDown()
}
