package user

import (
	"ecom-tiago/services/auth"
	"ecom-tiago/types"
	"ecom-tiago/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// docs: we are using repository pattern
// docs: inject store.go dependencies to this handler
type Handler struct {
	store types.UserStore;
}

// docs: create a new handler with the store dependency injected
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store};
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
	// w.Write([]byte("Signup success!"));

	// get JSON payload
	var payload types.RegisterUserPayload;
	if err:= utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err);
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email);
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s already exists", payload.Email));
		return;
	}

	// hash the password
	hashedPassword, err := auth.HashPassword(payload.Password);
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err);
		return;
	}

	// if it doesn't we create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	});
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err);
		return;
	}

	// utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"});
	utils.WriteJSON(w, http.StatusCreated, nil);
}