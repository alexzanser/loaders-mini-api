package handler

import (
	"context"
	"fmt"
	"loaders/internal/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type customerHandler struct {
	service *service.Service
}

func newCustomerHandler(service *service.Service) *customerHandler {
	return &customerHandler{service: service}
}

func (c *customerHandler) GetCustomer(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get customer at %s\n", req.URL.Path)

	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`{"error": "can't retreive username from context"}`), http.StatusBadRequest)
		return
	}

	passwd := req.PostFormValue("password")
	ct, err := c.service.GetCustomer(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "can't get customer": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	ld, err := c.service.GetLoadersList(req.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "can't get loaders list": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	rp := response{
		Username: ct.Username,
		Role:     "customer",
		Balance:  ct.Balance,
		Loaders:  ld,
	}

	renderResponse(w, rp)
}

func (c *customerHandler) GetCustomerTasks(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get customer tasks at %s\n", req.URL.Path)

	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`"error": "can't retrieve username from context"}`), http.StatusBadRequest)
		return
	}

	passwd := req.PostFormValue("password")
	ct, err := c.service.GetCustomer(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "can't get customer": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	if len(ct.Tasks) == 0 {
		http.Error(w, fmt.Sprintf(`"error": "task list is empty"}`), http.StatusBadRequest)
		return
	}

	rp := response{
		Username: ct.Username,
		Tasks:    ct.Tasks,
	}

	renderResponse(w, rp)
}

func (c *customerHandler) Start(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling start at %s\n", req.URL.Path)

	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`"error": "can't retrieve username from context"}`), http.StatusBadRequest)
		return
	}

	role, ok := req.Context().Value("role").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`"error": "can't retrieve role from context"}`), http.StatusBadRequest)
		return
	}

	if role != "customer" {
		http.Error(w, fmt.Sprintf(`"error": "acess denied"}`), http.StatusUnauthorized)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "invalid request": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	loadersStr := req.PostFormValue("loaders")
	loadersID := make([]int64, 0)
	for _, val := range strings.Split(loadersStr, ",") {
		v, err := strconv.Atoi(val)
		if err != nil {
			http.Error(w, fmt.Sprintf(`"error": "can't parse loaders ID"}`), http.StatusBadRequest)
			return
		}
		loadersID = append(loadersID, int64(v))
	}
	var rp response
	rp.Result, err = c.service.Start(context.TODO(), loadersID, username, "")
	if err != nil {
		rp.Err = err.Error()
	}

	renderResponse(w, rp)
}
