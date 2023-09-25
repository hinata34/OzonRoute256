package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-7/internal/app/user"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type serverUser struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
	Age  int64   `json:"age"`
}

type server struct {
	userRepo user.UserRepo
}

func (s *server) getUser(ctx context.Context, req *http.Request) ([]byte, int) {
	id, err := getUserID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}

	var user *user.User
	user, err = s.userRepo.GetById(ctx, id)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusInternalServerError
	}

	su := &serverUser{}
	su.ID = user.ID
	su.Name = &user.Name
	su.Age = user.Age

	data, err := json.Marshal(su)
	if err != nil {
		fmt.Errorf("can't marshal user with id: %d. Error: %s", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

func (s *server) getUsers(ctx context.Context, req *http.Request) ([][]byte, int) {
	var users []*user.User
	users, err := s.userRepo.List(ctx)
	if err != nil {
		fmt.Errorf("can't get all users: %s", err)
		return nil, http.StatusInternalServerError
	}
	result := [][]byte{}

	for _, user := range users {
		su := &serverUser{}
		su.ID = user.ID
		su.Name = &user.Name
		su.Age = user.Age

		data, err := json.Marshal(su)
		if err != nil {
			fmt.Errorf("can't marshal user with id: %d. Error: %s", su.ID, err)
			return nil, http.StatusInternalServerError
		}

		result = append(result, data)
	}

	return result, http.StatusOK
}

func (s *server) createUser(ctx context.Context, req *http.Request) (int64, int) {
	body, err := getUserData(req.Body)
	if err != nil {
		fmt.Errorf("can't parse request. Error: %s", err)
		return 0, http.StatusBadRequest
	}

	id, err := s.userRepo.Add(ctx, &user.User{Name: *body.Name, Age: body.Age})
	if err != nil {
		fmt.Errorf("can't create user. Error: %s", err)
		return 0, http.StatusInternalServerError
	}

	return id, http.StatusOK
}

func (s *server) updateUser(ctx context.Context, req *http.Request) (bool, int) {
	body, err := getUserData(req.Body)
	if err != nil {
		fmt.Errorf("can't parse request. Error: %s", err)
		return false, http.StatusBadRequest
	}

	ok, err := s.userRepo.Update(ctx, &user.User{ID: body.ID, Name: *body.Name, Age: body.Age})
	if err != nil {
		fmt.Errorf("can't update user. Error: %s", err)
		return false, http.StatusInternalServerError
	}

	if !ok {
		fmt.Errorf("can't find user. Error: %s", err)
		return false, http.StatusNotFound
	}

	return true, http.StatusOK
}

func CreateServer(ctx context.Context, ur user.UserRepo) *http.ServeMux {
	serv := server{
		userRepo: ur,
	}
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}

		switch req.Method {
		case http.MethodGet:
			data, status := serv.getUser(ctx, req)
			res.WriteHeader(status)
			res.Write(data)
		case http.MethodPost:
			id, status := serv.createUser(ctx, req)
			res.WriteHeader(status)
			value := map[string]int64{"id": id}
			json_data, _ := json.Marshal(value)
			res.Write(json_data)
		case http.MethodPut:
			_, status := serv.updateUser(ctx, req)
			res.WriteHeader(status)
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	return serveMux
}

func getUserID(reqUrl *url.URL) (int64, error) {
	idStr := reqUrl.Query().Get("id")
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getUserData(reader io.ReadCloser) (serverUser, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return serverUser{}, err
	}

	data := serverUser{}
	if err = json.Unmarshal(body, &data); err != nil {
		return data, err
	}

	return data, nil
}
