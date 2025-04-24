package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AliMumtaz001/GoTask/api/config"
	_ "github.com/lib/pq"
)

func Connect(dbConfig config.DBConfig) *sql.DB {
	// Create the connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName,
	)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully!")
	return db
}
