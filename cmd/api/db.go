package main

import (
    "database/sql"
    "log"
)

const (
    DBUser     = "your_username"
    DBPassword = "your_password"
    DBName     = "your_database_name"
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

