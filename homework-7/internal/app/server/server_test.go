package server

import (
	"bytes"
	"context"
	"encoding/json"
	"homework-7/internal/app/user"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_getUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()

		req, err := http.NewRequest(http.MethodGet, "user?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.repo.EXPECT().GetById(gomock.Any(), int64(id)).Return(&user.User{ID: 1, Name: "asd", Age: 25}, nil)

		_, status := s.serv.getUser(ctx, req)

		require.Equal(t, http.StatusOK, status)
	})
}

func Test_getUsers(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()

		req, err := http.NewRequest(http.MethodGet, "", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		s.repo.EXPECT().List(gomock.Any()).Return([]*user.User{
			{ID: 1, Name: "ha1", Age: 12},
			{ID: 2, Name: "ha2", Age: 13},
		}, nil)

		_, status := s.serv.getUsers(ctx, req)

		require.Equal(t, http.StatusOK, status)

	})
}

func Test_createUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()

		user_body := user.User{Name: "nice1", Age: 55}
		body, _ := json.Marshal(user_body)

		req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(body))
		require.NoError(t, err)

		s.repo.EXPECT().Add(gomock.Any(), &user.User{Name: "nice1", Age: 55}).Return(int64(1), nil)

		_, status := s.serv.createUser(ctx, req)

		require.Equal(t, http.StatusOK, status)

	})
}

func Test_updateUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()

		user_body := user.User{ID: 1, Name: "nice1", Age: 55}
		body, _ := json.Marshal(user_body)

		req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(body))
		require.NoError(t, err)

		s.repo.EXPECT().Update(gomock.Any(), &user.User{ID: 1, Name: "nice1", Age: 55}).Return(true, nil)

		_, status := s.serv.updateUser(ctx, req)

		require.Equal(t, http.StatusOK, status)
	})
}
