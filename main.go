package main

import (
	"fmt"
	"time"

	"github.com/AliMumtaz001/GoTask/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	start := time.Now()
	router.GET("/getData", getData)
	elapse0 := time.Since(start)
	fmt.Printf("Total execution took %s\n", elapse0)
	router.Run(":8080")
}

func getData(c *gin.Context) {
	result, err := utils.Analyzer("Test1.txt")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read file"})
		return
	}
	c.JSON(200, result)
}
