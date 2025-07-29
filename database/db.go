package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os" // Добавленный импорт
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}

	// Создание таблиц
	sqlBytes, err := os.ReadFile("database/init.sql")
	if err != nil {
		return err
	}

	sqlStmt := string(sqlBytes)
	_, err = DB.Exec(sqlStmt)
	return err
}
