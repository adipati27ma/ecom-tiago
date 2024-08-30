package types

import "time"

type User struct {
	ID				int    		`json:"id"`;
	FirstName string 		`json:"firstName"`;
	LastName 	string 		`json:"lastName"`;
	Email			string 		`json:"email"`;
	Password	string 		`json:"-"`;
	CreatedAt time.Time `json:"createdAt"`;
}

// docs: gunakan interface agar lebih mudah dalam pengujian/pengetesan
type UserStore interface {
	GetUserByEmail(email string) (*User, error);
	GetUserByID(id int) (*User, error);
	CreateUser(User) error;
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`;
	LastName 	string `json:"lastName"`;
	Email			string `json:"email"`;
	Password	string `json:"password"`;
}