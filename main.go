package main

import (
	"log"

	"github.com/AliMumtaz001/GoTask/api/config"
	"github.com/AliMumtaz001/GoTask/api/database"
	_ "github.com/AliMumtaz001/GoTask/docs" // Import generated Swagger docs
	"github.com/AliMumtaz001/GoTask/routes"
	"github.com/gin-gonic/gin"
)

// @title File Analyzer APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	envConfig := config.LoadEnv()

	port := envConfig.AppPort // Use AppPort for the application server
	if port == "" {
		port = "8000" // Default port
	}
	log.Printf("Using port: %s", port)

	db := database.Connect(envConfig)
	defer db.Close()

	r := gin.Default()

	routes.SetupRoutes(r, db, envConfig)

	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
