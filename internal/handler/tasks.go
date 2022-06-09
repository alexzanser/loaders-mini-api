package handler

import (
	"context"
	"encoding/binary"
	"fmt"
	"loaders/internal/service"
	"log"
	"net/http"
)

type tasksHandler struct {
	service *service.Service
}

func newTasksHandler(service *service.Service) *tasksHandler {
	return &tasksHandler{
		service: service,
	}
}

func (t *tasksHandler) GenerateRandomTasks(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling generate random tasks at %s\n", req.URL.Path)

	id, err := t.service.GenerateRandomTasks(context.TODO())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "can't generate new tasklist": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	b := make([]byte, 800)
	for _, val := range id {
		binary.LittleEndian.PutUint64(b, uint64(val))
	}
	renderResponse(w, response{Result: fmt.Sprintf("new tasklist created"),
		HTTPStatus: http.StatusCreated})
}
