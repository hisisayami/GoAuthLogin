package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
    DBUser     = "root"
    DBPassword = ""
    DBName     = "shajhya"
)

func GetDSN() string {
    return DBUser + ":" + DBPassword + "@/" + DBName
}

func connectToDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", GetDSN())
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return db, nil
}

