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
	r.POST("/result", func(c *gin.Context) { // Use a closure to pass db
		authentication.Upload(c, db)
	})
	r.GET(("/getdata"), authentication.GetResultHandler(db))
	r.Run(":8000")
}
