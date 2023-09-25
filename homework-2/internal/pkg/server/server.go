package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	headerKey = "hw-sum"
)

type Server struct{}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Create, headers: [%v]\n", r.Header)

	value := r.Header.Get(headerKey)
	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Sum: [%d]\n", intValue+5)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Update, body: [%s]\n", string(body))
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete")
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.RawQuery
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Get, query params: [%s]\n", key)
}
