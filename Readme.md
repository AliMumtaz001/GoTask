# GoTask

GoTask is a Go-based web application designed for analyzing text files and storing analysis results in a database. It follows a three-tier architecture, separating the presentation (HTTP handlers), business logic (services), and data access (repositories) layers. The application allows users to upload files, analyze their content for metrics such as word count, digit count, special characters, and more, and retrieve paginated results. Built with the Gin web framework, GoTask ensures scalability and maintainability.
Features

File Upload and Analysis: Upload text files and analyze content for various metrics (e.g., words, digits, special characters, lines, spaces, consonants, vowels, sentences, paragraphs, punctuation).
Three-Tier Architecture: Organized into handlers, services, and repositories for clear separation of concerns.
Pagination: Retrieve analysis results with pagination support for efficient data handling.
Database Integration: Store and retrieve analysis results using a SQL database.
Authentication: Secure endpoints with Bearer token authentication.
RESTful API: Expose endpoints for file upload and result retrieval using the Gin framework.

# Prerequisites
```json
Go (version 1.16 or later)
A SQL database (e.g., PostgreSQL, MySQL)
Git
Gin Gonic for HTTP routing
PostgreSQL driver or equivalent for your database
```

# Installation

Clone the Repository:
```json
git clone https://github.com/AliMumtaz001/GoTask.git
```
```json
cd GoTask
```


Install Dependencies:Ensure you have Go modules enabled, then run:
```json
go mod tidy
```


# Set Up the Database:

Create a database (e.g., in PostgreSQL):CREATE DATABASE gotask;

```json
Create the results table:CREATE TABLE results (
    id SERIAL PRIMARY KEY,
    words INT,
    digits INT,
    special_char INT,
    lines INT,
    spaces INT,
    consonants INT,
    vowels INT,
    sentences INT,
    paragraphs INT,
    punctuation INT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES employee(id)
);
```

```json
Create the employee table (for user authentication):CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);
```



Configure Environment Variables:Create a .env file in the project root with the following:
DATABASE_URL=your_database_connection_string
PORT=8080


Run the Application:
go run main.go

The server will start on http://localhost:8080 (or the port specified in .env).


# Usage

Upload a File:Use the /result endpoint to upload a text file for analysis. Ensure you include a valid Bearer token for authentication.
curl -X POST http://localhost:8080/result \
  -H "Authorization: Bearer your_token" \
  -F "file=@/path/to/your/file.txt"

# Response:
```json
{
    "status": 200,
    "message": "File uploaded and analyzed successfully",
    "file": "file.txt",
    "result": {
        "Words": 100,
        "Digits": 10,
        "SpecialChar": 5,
        "Lines": 20,
        "Spaces": 90,
        "Consonants": 200,
        "Vowels": 80,
        "Sentences": 10,
        "Paragraphs": 5,
        "Punctuation": 15
    }
}
```


# Retrieve Paginated Results:
Use the service layer’s GetPaginatedResults function to fetch analysis results for a user with pagination. This is typically accessed via an API endpoint (not shown in the provided code but can be implemented).

