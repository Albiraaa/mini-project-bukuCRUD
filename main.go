package main

import (
	"os"
	"projek/controllers"
	"projek/database"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set")
	}

	dsn = dsn + "?sslmode=require"

	database.ConnectDatabase(dsn)

	// Auto migrate tables
	database.AutoMigrate()

	r := gin.Default()

	// CATEGORY ROUTES
	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)
	r.PUT("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)

	// BOOK ROUTES
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
