package handler

import (
	"context"
	"fmt"
	"loaders/internal/service"
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
	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive username from context"), http.StatusBadRequest)
		return
	}
	passwd := req.PostFormValue("password")
	ct, err := c.service.GetCustomer(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when get customer :%v", err), http.StatusInternalServerError)
		return
	}
	ld, err := c.service.GetLoadersList(req.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("error when get loaders list :%v", err), http.StatusInternalServerError)
		return
	}
	rp := Response {
		Username:	ct.Username,
		Role: 		"customer",	
		Balance: 	ct.Balance,
		Loaders: 	ld,
	}
	renderResponse(w, rp)
}

func (c *customerHandler) GetCustomerTasks(w http.ResponseWriter, req *http.Request) {
	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive username from context"), http.StatusBadRequest)
		return
	}
	passwd := req.PostFormValue("password")
	ct, err := c.service.GetCustomer(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when get customer :%v", err), http.StatusInternalServerError)
		return
	}

	if len(ct.Tasks) == 0 {
		http.Error(w, fmt.Sprintf("tasklist is empty"), http.StatusInternalServerError)
		return
	}
	rp := Response {
		Username:	ct.Username,
		Tasks: 		ct.Tasks,
	}
	renderResponse(w, rp)
}

func (c *customerHandler) Start(w http.ResponseWriter, req *http.Request) {
	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive username from context"), http.StatusBadRequest)
		return
	}
	role, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive role from context"), http.StatusBadRequest)
		return
	}

	if role != "customer" {
			http.Error(w, fmt.Sprintf(`{"\start is only for customers":`), http.StatusUnauthorized)
			return
	}

	loadersStr := req.PostFormValue("loaders")
	loadersID := make([]int64, 0)
	for _, val := range strings.Split(loadersStr, ",") {
		v, err := strconv.Atoi(val)
		if err != nil {
			http.Error(w, fmt.Sprintf("loader id should be int"), http.StatusBadRequest)
			return
		}
		loadersID = append(loadersID, int64(v))
	}
	
	c.service.Start(context.TODO(), loadersID, username, "")
	renderResponse(w, rp)
}