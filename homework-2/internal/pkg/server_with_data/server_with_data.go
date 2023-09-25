package server_with_data

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const (
	queryParamKey = "id"
)

type ServerWithData struct {
	Data map[uint32]string
}

type request struct {
	Key   *uint32 `json:"id"`
	Value *string `json:"value"`
}

func receiveBody(w http.ResponseWriter, r *http.Request) (unmarshalled request, error bool) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error = true
		return
	}

	if err := json.Unmarshal(body, &unmarshalled); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error = true
		return
	}

	if unmarshalled.Key == nil || unmarshalled.Value == nil {
		w.WriteHeader(http.StatusBadRequest)
		error = true
		return
	}

	return
}

func (s *ServerWithData) Create(w http.ResponseWriter, r *http.Request) {
	unmarshalled, err := receiveBody(w, r)
	if err {
		return
	}

	if _, ok := s.Data[*unmarshalled.Key]; ok {
		w.WriteHeader(http.StatusConflict)
		return
	}

	s.Data[*unmarshalled.Key] = *unmarshalled.Value
}

func (s *ServerWithData) Update(w http.ResponseWriter, r *http.Request) {
	unmarshalled, err := receiveBody(w, r)
	if err {
		return
	}

	if _, ok := s.Data[*unmarshalled.Key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.Data[*unmarshalled.Key] = *unmarshalled.Value
}

func (s *ServerWithData) Delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get(queryParamKey)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intKey, err := strconv.ParseUint(key, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.Data[uint32(intKey)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.Data, uint32(intKey))
}

func (s *ServerWithData) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get(queryParamKey)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intKey, err := strconv.ParseUint(key, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := s.Data[uint32(intKey)]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := w.Write([]byte(value)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
