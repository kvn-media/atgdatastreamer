package database

import (
    "database/sql"
    "log"
)

var db *sql.DB

// InitDB inisialisasi koneksi ke SQLite
func InitDB(dbPath string) (*sql.DB, error) {
    var err error
    db, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    log.Println("Connected to the database")
    return db, nil
}

// CloseDB menutup koneksi database
func CloseDB() {
    if db != nil {
        db.Close()
    }
    log.Println("Connection to the database closed")
}