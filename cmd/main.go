package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/memojito/igapi/handler"
	"github.com/memojito/igapi/middleware"
	"github.com/memojito/igapi/storage"
	"github.com/memojito/igapi/types"
	"github.com/memojito/igapi/utils"
)

type Server struct {
	listenAddr string
	storage    storage.Storage
}

func main() {
	url := os.Getenv("POSTGRESQL_URL")
	if url == "" {
		log.Fatal("POSTGRESQL_URL not found")
	}

	ps, err := storage.New(url)
	if err != nil {
		log.Fatalf("connection failed %v", err)
		return
	}

	err = ps.Conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("ping failed %v", err)
	}

	h := handler.New(ps)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", makeHandler(h.GetAllUsers))
	mux.HandleFunc("GET /users/{id}", makeHandler(h.GetUser))
	mux.HandleFunc("POST /users", makeHandler(h.AddUser))
	if err := http.ListenAndServe(":8080", middleware.Logging(mux)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHandler(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var e types.APIError
			if errors.As(err, &e) {
				log.Printf("API error: %s \n Status: %d", e.Msg, e.Status)
				err := utils.WriteJSON(w, e.Status, e)
				if err != nil {
					log.Printf("could not write ERROR response: %v", err)
				}
			}
		}
	}
}
