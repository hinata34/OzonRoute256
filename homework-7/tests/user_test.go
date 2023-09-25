//go:build integration
// +build integration

package tests

import (
	"context"
	"homework-7/internal/app/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_createUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)
	t.Run("successs", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		_, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)
		_, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})

		assert.NoError(t, err)
	})
}

func Test_getUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("successs", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		id, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})

		user, err := userRepo.GetById(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, "ivanov", user.Name)
		assert.Equal(t, int64(24), user.Age)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		id, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})

		_, err = userRepo.GetById(ctx, id+1)

		assert.Error(t, err)
	})

}

func Test_getUsers(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("successs", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		_, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})
		_, err = userRepo.Add(ctx, &user.User{Name: "petrov", Age: 13})
		users, err := userRepo.List(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "ivanov", users[0].Name)
		assert.Equal(t, int64(24), users[0].Age)

		assert.Equal(t, "petrov", users[1].Name)
		assert.Equal(t, int64(13), users[1].Age)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		_, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})
		_, err = userRepo.Add(ctx, &user.User{Name: "petrov", Age: 13})
		users, err := userRepo.List(ctx)

		assert.NoError(t, err)
		assert.NotEqual(t, "", users[0].Name)
		assert.NotEqual(t, int64(0), users[0].Age)

		assert.NotEqual(t, "", users[1].Name)
		assert.NotEqual(t, int64(0), users[1].Age)
	})
}

func Test_updateUser(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)
	t.Run("successs", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		id, err := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})
		ok, err := userRepo.Update(ctx, &user.User{ID: id, Name: "petrov", Age: 13})

		require.NoError(t, err)
		require.Equal(t, true, ok)

		user, err := userRepo.GetById(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, "petrov", user.Name)
		assert.Equal(t, int64(13), user.Age)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		Db.SetUp(t)
		defer Db.TearDown()

		userRepo := user.NewUsers(Db.DB)

		id, _ := userRepo.Add(ctx, &user.User{Name: "ivanov", Age: 24})
		ok, _ := userRepo.Update(ctx, &user.User{ID: id + 1, Name: "petrov", Age: 13})

		require.NotEqual(t, true, ok)

		ok, _ = userRepo.Update(ctx, &user.User{ID: id, Name: "petrov", Age: 13})
		user, _ := userRepo.GetById(ctx, id)

		assert.NotEqual(t, "ivanov", user.Name)
		assert.NotEqual(t, int64(24), user.Age)
	})
}
