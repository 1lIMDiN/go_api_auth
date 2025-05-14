package database

import (
	"database/sql"
	"fmt"
	"os"

	_"modernc.org/sqlite"
)

var DB *sql.DB

func Init(dbFile string) error {
	// Проверка существования БД
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	// Открываем БД
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Если не существовал - выполняем команды SchemaDB
	if install {
		_, err := db.Exec(SchemaDB)
		if err != nil {
			return fmt.Errorf("failed to create schema: %w", err)
		}
	}

	DB = db

	return nil
}
