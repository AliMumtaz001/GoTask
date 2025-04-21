package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv" // Add this import for strconv.Atoi

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validate that all required environment variables are set
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Missing required environment variables")
	}

	// Convert port to int for the connection string
	portInt, err := strconv.Atoi(port) // Use strconv.Atoi to convert string to int
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	// Create the connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	fmt.Println("Connected to PostgreSQL successfully!")

	return db
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error for result:", err)
		log.Fatal(err)

	}
}
