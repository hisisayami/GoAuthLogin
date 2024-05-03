package models

import (
	//"errors"
	"time"

	//"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	//Password  string    `json:"password"`
	Access    bool      `json:"access"`
	Phone     int       `json:"phone"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type LoginResponse struct {
	User    *User  `json:"user,omitempty"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}

// func (u *User) PasswordMatches(plainText string) (bool, error) {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
// 			//invalid password
// 			return false, nil
// 		default:
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }
