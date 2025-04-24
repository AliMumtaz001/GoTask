package routes

import (
    "database/sql"

    "github.com/AliMumtaz001/GoTask/authentication"
    "github.com/AliMumtaz001/GoTask/handlers"
    "github.com/AliMumtaz001/GoTask/api/config"
    "github.com/gin-gonic/gin"
    swagFiles "github.com/swaggo/files"        // Correct alias for Swagger files
    ginSwagger "github.com/swaggo/gin-swagger" // Gin Swagger middleware
)

func SetupRoutes(r *gin.Engine, db *sql.DB, envConfig config.DBConfig) {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

    // Public routes (no authentication)
    r.POST("/login", func(c *gin.Context) {
        authentication.Login(c, db, envConfig)
    })
    r.POST("/signup", func(c *gin.Context) {
        authentication.Signup(c, db, envConfig)
    })

    // Protected routes (require authentication)
    protected := r.Group("/")
    protected.Use(authentication.Auths())
    {
        protected.POST("/result", func(c *gin.Context) {
            handlers.Upload(c, db)
        })
        protected.GET("/getdata", handlers.GetResultHandler(db))
    }
}