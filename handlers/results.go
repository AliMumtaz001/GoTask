package handlers

import (
	"database/sql"
	"net/http"

	"github.com/AliMumtaz001/GoTask/api/repositories"
	"github.com/AliMumtaz001/GoTask/services"

	"github.com/gin-gonic/gin"
)

// Result godoc
// @Summary      Upload file for analysis and get result
// @Description  Upload a file, analyze its content, return the analysis result as JSON, and store the result in the database
// @Tags         file
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload for analysis"
// @Success      200     "File uploaded and analyzed successfully"
// @Failure      400     "Bad Request"
// @Failure      401     "Unauthorized"
// @Failure      403     "Forbidden"
// @Failure      500     "Internal Server Error"
// @Router       /result [post]
// @Security     BearerAuth
func Upload(c *gin.Context, db *sql.DB) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
		return
	}

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	repo := &repositories.ResultsRepository{DB: db}
	service := services.NewResultsService(repo)

	result, err := service.ProcessFile(c, file, email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "File uploaded and analyzed successfully",
		"file":    file.Filename,
		"result":  result,
	})
}
