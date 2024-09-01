package user

import (
	"database/sql"
	"ecom-tiago/types"
	"fmt"
)

// docs: using repository pattern
// this is "store" repository
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// docs: ? operator is used to prevent SQL Injection, is a placeholder, and the value is passed as a second argument
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	// docs: kalo pusing, baca aja penjelasan dari tiap Method yg digunakan :D
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// docs: if the user is not found, return an error
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// generated by Copilot (touched)
func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// generated by Copilot (touched)
func (s *Store) CreateUser(user types.User) error {
	// docs: ? operator is used to prevent SQL Injection, is a placeholder, and the value is passed as a second argument
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// docs: scanRowIntoUser is a helper function to scan the row into a User struct
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	// docs: scan the row and store values into the user struct
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
