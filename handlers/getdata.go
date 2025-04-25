package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AliMumtaz001/GoTask/services"
	"github.com/gin-gonic/gin"
)

// GetResultHandler godoc
// @Summary      Get analysis results for a user
// @Description  Retrieve paginated analysis results for a specific user based on user_id
// @Tags         file
// @Accept       json
// @Produce      json
// @Param        user_id   query  string  true  "User ID to fetch results for"
// @Param        page      query  int     false "Page number (default: 1)"
// @Param        page_size query  int     false "Number of results per page (default: 10)"
// @Success      200  {object}  map[string]interface{}  "Successfully retrieved results"
// @Failure      400  {object}  map[string]interface{}  "Bad Request"
// @Failure      401  {object}  map[string]interface{}  "Unauthorized"
// @Failure      404  {object}  map[string]interface{}  "Not Found"
// @Router       /getdata [get]
// @Security     BearerAuth
func GetResultHandler(service *services.ResultsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id query parameter"})
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("page_size")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}

		jsonData, totalRecords, err := service.GetPaginatedResults(userID, page, pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or server error", "details": err.Error()})
			return
		}

		if jsonData == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "No results found for this user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id":       userID,
			"page":          page,
			"page_size":     pageSize,
			"total_records": totalRecords,
			"total_pages":   (totalRecords + pageSize - 1) / pageSize,
			"results":       json.RawMessage(jsonData),
		})
	}
}
