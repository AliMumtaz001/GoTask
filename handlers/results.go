package handlers

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

	resultprocess "github.com/AliMumtaz001/GoTask/result"
	"github.com/AliMumtaz001/GoTask/utils"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, db *sql.DB) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
		return
	}

	if err := os.MkdirAll("upload", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
		return
	}

	Path := filepath.Join("upload", file.Filename)
	err = c.SaveUploadedFile(file, Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	content, err := os.ReadFile(Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the uploaded file"})
		return
	}

	// Get the email from the context (set by Auths middleware)
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Query the employeedata table to get the user_id
	var userID int
	query := `SELECT id FROM employee WHERE email = $1`
	err = db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID", "details": err.Error()})
		return
	}

	data := string(content)
	res := utils.CombineFunc(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "File uploaded and analyzed successfully",
		"file":    file.Filename,
		"result":  res,
	})

	err = resultprocess.SaveResult(db, res, userID) // Pass userID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save result to database", "details": err.Error()})
	}
}
