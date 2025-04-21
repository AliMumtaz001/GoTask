// package main

// import (
// 	// "github.com/AliMumtaz001/GoTask/utils"
// 	"github.com/AliMumtaz001/GoTask/authentication"
// 	"github.com/AliMumtaz001/GoTask/database"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// router := gin.Default()
// 	// start := time.Now()
// 	// router.GET("/getData", getData)
// 	// elapse0 := time.Since(start)
// 	// fmt.Printf("Total execution took %s\n", elapse0)
// 	// router.Run(":8080")
// 	r := gin.Default()

// 	r.POST("/login", authentication.Login)
// 	r.POST("/signup", authentication.Signup)
// 	r.Use(authentication.Auths())
// 	r.GET("/result", authentication.SaveResult)

// 	r.Run(":8080")
// 	database.Connect()
// }

// // func getData(c *gin.Context) {
// // 	result, err := utils.Analyzer("Test.txt")
// // 	if err != nil {
// // 		c.JSON(500, gin.H{"error": "Failed to read file"})
// // 		return
// // 	}
// // 	c.JSON(200, result)
// // }

package main

import (
	"github.com/AliMumtaz001/GoTask/authentication"
	"github.com/AliMumtaz001/GoTask/database"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.Connect()
	defer db.Close() 


	r := gin.Default()

	// routes
	r.POST("/login", authentication.Login)
	r.POST("/signup", authentication.Signup)
	r.Use(authentication.Auths())
	r.POST("/result", authentication.Upload)
	r.Run(":8080")
}
