package handler

import (
	"fmt"
	"loaders/internal/service"
	"log"
	"net/http"
)

type loaderHandler struct {
	service *service.Service
}

func newLoaderHandler(service *service.Service) *loaderHandler {
	return &loaderHandler{service: service}
}

func (c *loaderHandler) GetLoader(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get loader at %s\n", req.URL.Path)

	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`"error": "can't retrieve username from context"}`), http.StatusBadRequest)
		return
	}

	passwd := req.PostFormValue("password")
	ld, err := c.service.GetLoader(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "can't get loader": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	rp := response{
		Username: ld.Username,
		Role:     "loader",
		Balance:  ld.Balance,
		Salary:   ld.Salary,
		Weight:   ld.MaxWeight,
		Drunk:    ld.Drunk,
		Fatigue:  ld.Fatigue,
	}

	renderResponse(w, rp)
}

func (c *loaderHandler) GetLoaderTasks(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get loader at %s\n", req.URL.Path)

	username, ok := req.Context().Value("username").(string)
	if ok == false {
		http.Error(w, fmt.Sprintf(`"error": "can't retrieve username from context"}`), http.StatusBadRequest)
		return
	}

	passwd := req.PostFormValue("password")
	ld, err := c.service.GetLoader(req.Context(), username, passwd)
	if err != nil {
		http.Error(w, fmt.Sprintf(`"error": "can't get loader": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	rp := response{
		Username: ld.Username,
		Role:     "loader",
		Tasks:    ld.CompletedTasks,
	}

	renderResponse(w, rp)
}
