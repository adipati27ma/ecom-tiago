package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{};
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// docs: this is subRouter, so the path is /api/v1/login
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost);
	router.HandleFunc("/signup", h.handleSignup).Methods(http.MethodPost);
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login success!"));
}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup success!"));
}