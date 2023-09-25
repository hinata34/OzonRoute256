package main

import (
	"fmt"
	"homework/internal/pkg/server"
	"homework/internal/pkg/server_with_data"
	"log"
	"net/http"
)

const (
	portServer         = ":9000"
	portServerWithData = ":9001"
)

func main() {
	server := server.Server{}

	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.Get(w, r)
		case http.MethodDelete:
			server.Delete(w, r)
		case http.MethodPost:
			server.Create(w, r)
		case http.MethodPut:
			server.Update(w, r)
		default:
			fmt.Printf("unsupported method: [%s]\n", r.Method)
		}
	})

	serverWithData := server_with_data.ServerWithData{
		Data: map[uint32]string{},
	}

	muxServerWithData := http.NewServeMux()
	muxServerWithData.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			serverWithData.Get(w, r)
		case http.MethodDelete:
			serverWithData.Delete(w, r)
		case http.MethodPost:
			serverWithData.Create(w, r)
		case http.MethodPut:
			serverWithData.Update(w, r)
		default:
			fmt.Printf("unsupported method: [%s]\n", r.Method)
		}
	})

	go func() {
		if err := http.ListenAndServe(portServerWithData, muxServerWithData); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(portServer, muxServer); err != nil {
		log.Fatal(err)
	}
}
