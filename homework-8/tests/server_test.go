//go:build integration
// +build integration

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homework-7/internal/app/user"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getUserServer(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		UserServer.SetUp(t)
		defer UserServer.TearDown()

		userData := user.User{Name: "lala", Age: 3}
		jsonData, _ := json.Marshal(userData)
		res, err := http.Post("http://localhost:1337/user", "application/json", bytes.NewBuffer(jsonData))

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)

		data, err := io.ReadAll(res.Body)
		us := user.User{}
		err = json.Unmarshal(data, &us)

		requestUrl := fmt.Sprintf("http://localhost:1337/user?id=%d", us.ID)
		res, err = http.Get(requestUrl)

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
