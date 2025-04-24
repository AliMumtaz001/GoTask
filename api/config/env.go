package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func LoadEnv() DBConfig {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve environment variables
    config := DBConfig{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
    }

    // Validate that all required environment variables are set
    if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" {
        log.Fatal("Missing required environment variables")
    }

    return config
}