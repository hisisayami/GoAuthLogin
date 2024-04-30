package dbrepo

import (
	"database/sql"
	//"time"
)

type SqlDBRepo struct {
	DB *sql.DB
}

//const dbTimeout = time.Second * 3

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