package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"loaders/internal/handler"
	"net/http"
	"loaders/internal/repository"
	"loaders/internal/service"
	"loaders/pkg/postgres"
)

const connstr = "postgres://user_go:8956_go@db:5432/loaders"

func main() {
	dbPool, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatalf("Error connecting database: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.NewRepository(dbPool, nil)

	service := service.NewService(repo)

	handler := handler.NewHandler(service)
	r := chi.NewRouter()
	r.Mount("/", handler.Routes())

	log.Fatal(http.ListenAndServe(":8080", r))
}
