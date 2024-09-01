package user

import (
	"ecom-tiago/configs"
	"ecom-tiago/services/auth"
	"ecom-tiago/types"
	"ecom-tiago/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	// w.Write([]byte("Login success!"));

	// get JSON payload
	var payload types.LoginUserPayload;
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err);
		return;
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors);
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors));
		return;
	}

	// docs: get the user by email
	u, err := h.store.GetUserByEmail(payload.Email);
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s not found, or invalid password", payload.Email));
		return;
	}

	// docs: compare the password
	if !auth.ComparePasswords(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with email %s not found, or invalid password", payload.Email));
		return;
	}

	// docs: create JWT token
	secret := []byte(configs.Envs.JWTSecret);
	token, err := auth.CreateJWT([]byte(secret), u.ID);
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err);
		return;
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Login success!", "token": token});
}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Signup success!"));

	// get JSON payload
	var payload types.RegisterUserPayload;
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err);
		return;
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors);
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors));
		return;
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email);
	// docs: if no error, then the user exists
	if err == nil {
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