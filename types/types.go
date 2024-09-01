package types

import "time"

type User struct {
	// docs: the tags `json:...` used for encoding and decoding JSON
	ID				int    		`json:"id"`;
	FirstName string 		`json:"firstName"`;
	LastName 	string 		`json:"lastName"`;
	Email			string 		`json:"email"`;
	Password	string 		`json:"-"`; // docs: json:"-" agar tidak di tampilkan di response
	CreatedAt time.Time `json:"createdAt"`;
}

// docs: gunakan interface agar lebih mudah dalam pengujian/pengetesan
type UserStore interface {
	GetUserByEmail(email string) (*User, error);
	GetUserByID(id int) (*User, error);
	CreateUser(User) error;
}

type RegisterUserPayload struct {
	// docs: validate using go-playground/validator
	FirstName string `json:"firstName" validate:"required"`;
	LastName 	string `json:"lastName" validate:"required"`;
	Email			string `json:"email" validate:"required,email"`;
	Password	string `json:"password" validate:"required,min=6,max=100"`;
}