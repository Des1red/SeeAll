// initilize db
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() {
	if err := initdb("seeall.db"); err != nil {
		panic(err)
	}
}

func initdb(path string) error {
	var err error

	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	DB.Exec("PRAGMA journal_mode=WAL;")

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	if err := os.Chmod(path, 0600); err != nil {
		log.Printf("Warning: failed to set file permissions: %v", err)
	}

	return applySchema()
}

const schemaDir = "internal/database/schema/"

func applySchema() error {
	// Run schema files
	if err := execSQLFile(schemaDir + "001_tokens.sql"); err != nil {
		log.Fatalf("Failed to execute 001_tokens.sql: %v", err)
	}
	if err := execSQLFile(schemaDir + "002_metrics.sql"); err != nil {
		log.Fatalf("Failed to execute 002_metrics.sql: %v", err)
	}
	return nil
}

func execSQLFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read 2 file %s: %w", filename, err)
	}

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := DB.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query in %s: %w", filename, err)
		}
	}
	return nil
}
