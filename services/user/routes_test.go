package user

import (
	"bytes"
	"ecom-tiago/types"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServicesHandlers(t *testing.T) {
	userStore := &mockUserStore{};
	handler := NewHandler(userStore);

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Tiago",
			LastName:  "Lopez",
			Email:     "invalid",
			Password:  "password",
		};
		// docs: Marshal returns the JSON encoding of the arguments.
		marshalled, _ := json.Marshal(payload);
		
		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(marshalled));
		if err != nil {
			t.Fatal(err);
		}

		// docs: returns a new ResponseRecorder.
		rr := httptest.NewRecorder();
		router := mux.NewRouter();

		router.HandleFunc("/signup", handler.handleSignup);
		// make the request
		router.ServeHTTP(rr, req);

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code);
		}
	});

	t.Run("Should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Tiago",
			LastName:  "Lopes",
			Email:     "valid@gmail.com",
			Password:  "password",
		};
		// docs: Marshal returns the JSON encoding of the arguments.
		marshalled, _ := json.Marshal(payload);
		
		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(marshalled));
		if err != nil {
			t.Fatal(err);
		}

		// docs: returns a new ResponseRecorder.
		rr := httptest.NewRecorder();
		router := mux.NewRouter();

		router.HandleFunc("/signup", handler.handleSignup);
		// make the request
		router.ServeHTTP(rr, req);

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code);
		}
	});
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found");
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil;
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil;
}