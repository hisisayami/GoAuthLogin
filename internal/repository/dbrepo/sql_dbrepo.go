package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"goauthlogin/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Access    bool      `json:"access"`
	Phone     int       `json:"phone"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type SqlDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 10

func (m *SqlDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *SqlDBRepo) CreateUserTable() error {
	_, err := m.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            email VARCHAR(255) NOT NULL,
            username VARCHAR(255) NOT NULL,
            access INT NOT NULL,
            phone VARCHAR(20) NOT NULL,
            password VARCHAR(255) NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	return nil
}

func (m *SqlDBRepo) GetUserName(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, FirstName, LastName, UserName, Password, Access, Phone 
	from users where username = ?`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, username)

	err := row.Err()
	if err != nil {
		fmt.Println("failed to query the email", err)
		return nil, err
	}

	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Password,
		&user.Access,
		&user.Phone,
	)
	if err != nil {
		_ = models.LoginResponse{
			User:    nil,
			Message: "Authentication unsuccessful",
			Success: false,
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser inserts a new user record into the database
func (m *SqlDBRepo) CreateUser(user *models.User) error {
	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//Set the created_at and updated_at timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Execute the SQL query to insert the user into the database
	_, err = m.DB.Exec(`
        INSERT INTO users (firstname, lastname, username, password, access, phone)
        VALUES (?, ?, ?, ?, ?, ?)
    `,  user.FirstName, user.LastName, user.UserName, user.Password, user.Access, user.Phone)
	if err != nil {
		return err
	}

	return nil
}
