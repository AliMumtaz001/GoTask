package handlers

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AliMumtaz001/GoTask/repositories"
	"github.com/AliMumtaz001/GoTask/services"
	"github.com/AliMumtaz001/GoTask/utils"

	"github.com/gin-gonic/gin"
)

// Result godoc
// @Summary      Upload file for analysis and get result
// @Description  Upload a file, analyze its content, return the analysis result as JSON, and store the result in the database
// @Tags         file
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload for analysis"
// @Success      200  {object}  map[string]interface{}  "File uploaded and analyzed successfully"
// @Failure      400  {object}  map[string]interface{}  "Bad Request"
// @Failure      401  {object}  map[string]interface{}  "Unauthorized"
// @Failure      403  {object}  map[string]interface{}  "Forbidden"
// @Failure      500  {object}  map[string]interface{}  "Internal Server Error"
// @Router       /result [post]
// @Security     BearerAuth
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

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var userID int
	query := `SELECT id FROM employee WHERE email = $1`
	err = db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID", "details": err.Error()})
		return
	}

	data := string(content)
	res := utils.CombineFunc(data)

	// Initialize the service with the repository
	repo := &repositories.ResultsRepository{DB: db}
	service := services.NewResultsService(repo)

	// Save the result using the service
	err = service.SaveResult(res, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save result to database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "File uploaded and analyzed successfully",
		"file":    file.Filename,
		"result":  res,
	})
}