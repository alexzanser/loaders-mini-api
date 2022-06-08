package handler

import (
	"loaders/internal/service"
	"net/http"
	"fmt"
)

type loaderHandler struct {
	service *service.Service
}

func newLoaderHandler(service *service.Service) *loaderHandler {
	return &loaderHandler{service: service}
}

func (c *loaderHandler) GetLoader(w http.ResponseWriter, req *http.Request) {
	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive username from context"), http.StatusBadRequest)
		return
	}
	passwd := req.PostFormValue("password")
	ld, err := c.service.GetLoader(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when get loader :%v", err), http.StatusInternalServerError)
		return
	}
	rp := Response {
		Username:	ld.Username,
		Role: 		"loader",
		Balance: 	ld.Balance,
		Weight: 	ld.MaxWeight,
		Drunk: 		ld.Drunk,
		Fatigue: 	ld.Fatigue,	
	}

	renderResponse(w, rp)
}

func (c *loaderHandler) GetLoaderTasks(w http.ResponseWriter, req *http.Request) {
	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf("can't retreive username from context"), http.StatusBadRequest)
		return
	}
	passwd := req.PostFormValue("password")
	ld, err := c.service.GetLoader(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when get loader :%v", err), http.StatusInternalServerError)
		return
	}
	rp := Response {
		Username:	ld.Username,
		Role: 		"loader",
		Tasks: 		ld.CompletedTasks,
	}

	renderResponse(w, rp)
}