package sqlite

import (
    "database/sql"
    "log"
    "os"
    "path/filepath"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    path := getDataFolderPath()
    dbPath := filepath.Join(path, "todo.db")

    var err error

    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatalf("failed to connect to sqlite: %v", err)
    }

    createTodoTable()
}

func createTodoTable() {
    createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "title" TEXT,
        "completed" INTEGER
    );`

    statement, err := DB.Prepare(createTableSQL)
    if err != nil {
        log.Fatalf("failed to prepare create table statement: %v", err)
    }
    statement.Exec()
}

func getDataFolderPath() string {
    path := os.Getenv("DATA_FOLDER")
    if path == "" {
        path = "."
    }
    err := os.MkdirAll(path, os.ModePerm)
    if err != nil {
        log.Fatalf("failed to create data folder: %v", err)
    }
    return path
}
