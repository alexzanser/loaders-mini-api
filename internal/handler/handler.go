package handler

import (
	"github.com/go-chi/chi/v5"
	"loaders/internal/service"
	"net/http"
)

type handler struct {
	customerHandler			*customerHandler
	registerHandler			*registerHandler
	loginHandler			*loginHandler
	authorizationHandler	*authorizationHandler
	tasksHandler			*tasksHandler
	loaderHandler			*loaderHandler
}

func NewHandler(service *service.Service) *handler {
	return &handler{
		customerHandler:		newCustomerHandler(service),
		loaderHandler:			newLoaderHandler(service),
		registerHandler: 		newRegisterHandler(service),
		loginHandler:			newLoginHandler(service),
		authorizationHandler: 	newAuthorizationHandler(service),
		tasksHandler: 			newTasksHandler(service),
	}
}

func (h *handler) GetUser(w http.ResponseWriter, req *http.Request ) {
	role, _ := req.Context().Value("role").(string)
	if role == "loader" {
		h.loaderHandler.GetLoader(w, req)
		return
	}
	if role == "customer" {
		h.customerHandler.GetCustomer(w, req)
		return
	}
}

func (h *handler) GetTasks(w http.ResponseWriter, req *http.Request ) {
	role, _ := req.Context().Value("role").(string)
	if role == "loader" {
		h.loaderHandler.GetLoaderTasks(w, req)
		return
	}
	if role == "customer" {
		h.customerHandler.GetCustomerTasks(w, req)
		return
	}
}


func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/register", h.registerHandler.Register)
		r.Post("/login", h.loginHandler.Login)
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", h.tasksHandler.GenerateRandomTasks)
		r.With(h.authorizationHandler.Authorize).Get("/", h.GetTasks)
	})
	
	r.Route("/me", func(r chi.Router) {
		r.With(h.authorizationHandler.Authorize).Get("/", h.GetUser)
	})

	r.Route("/start", func(r chi.Router) {
		r.With(h.authorizationHandler.Authorize).Get("/", h.customerHandler.Start)
	})

	return r
}
