package handlers

import (
	"net/http"
	"strconv"

	model "github.com/AliMumtaz001/GoTask/models"
	"github.com/AliMumtaz001/GoTask/services"
	"github.com/gin-gonic/gin"
)

func SaveResultHandler(service *services.ResultsService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var result model.Multiples
        if err := c.ShouldBindJSON(&result); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
            return
        }

        userIDStr := c.Query("user_id")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil || userID <= 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
            return
        }

        err = service.SaveResult(result, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Result saved successfully"})
    }
}