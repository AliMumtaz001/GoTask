package main

import (
	"github.com/AliMumtaz001/GoTask/authentication"
	"github.com/AliMumtaz001/GoTask/database"
	handlers "github.com/AliMumtaz001/GoTask/handlers"

	_ "github.com/AliMumtaz001/GoTask/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Swagger files
	ginSwagger "github.com/swaggo/gin-swagger" // Gin Swagger middleware
)

// @title File Analyzer APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

//@in header
//@name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := database.Connect()
	defer db.Close()

	r := gin.Default()

	// Swagger route (no authentication required)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes (no authentication)
	r.POST("/login", authentication.Login)
	r.POST("/signup", authentication.Signup)

	// Protected routes (require authentication)
	protected := r.Group("/")
	protected.Use(authentication.Auths())
	{
		protected.POST("/result", func(c *gin.Context) {
			handlers.Upload(c, db)
		})
		protected.GET("/getdata", handlers.GetResultHandler(db))
	}

	r.Run(":8000")
}
