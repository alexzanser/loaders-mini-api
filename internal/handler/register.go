package handler

import (
	"context"
	"fmt"
	"loaders/internal/models"
	"loaders/internal/service"
	"log"
	"net/http"
)

type registerHandler struct {
	service *service.Service
}

func newRegisterHandler(service *service.Service) *registerHandler {
	return &registerHandler{service: service}
}

func (a *registerHandler) Register(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling register at %s\n", req.URL.Path)

	err := req.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't register user": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	role := req.PostFormValue("role")
	if role != "customer" && role != "loader" {
		http.Error(w, fmt.Sprintf(`{"error": "can't register user": "invalid role"}`), http.StatusBadRequest)
		return
	}

	var id int64
	if role == "customer" {
		ct := models.NewCustomer()
		if err = formToCustomerRegister(ct, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err = a.service.CreateCustomer(context.TODO(), ct)
	}

	if role == "loader" {
		ld := models.NewLoader()
		if err = formToLoaderRegister(ld, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err = a.service.CreateLoader(context.TODO(), ld)
	}

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't register user": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	renderResponse(w, response{Result: fmt.Sprintf("new user with role %s and id %d created", role, id),
		HTTPStatus: http.StatusCreated})
}
