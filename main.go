package main

import (
	"github.com/AliMumtaz001/GoTask/utils"
	"github.com/gin-gonic/gin"
	"github.com/AliMumtaz001/GoTask/database"
)

func main() {
	// router := gin.Default()
	// start := time.Now()
	// router.GET("/getData", getData)
	// elapse0 := time.Since(start)
	// fmt.Printf("Total execution took %s\n", elapse0)
	// router.Run(":8080")

	database.Connect()
}

func getData(c *gin.Context) {
	result, err := utils.Analyzer("Test.txt")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read file"})
		return
	}
	c.JSON(200, result)
}
