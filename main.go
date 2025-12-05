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
	r.GET("/kategori", controllers.GetCategories)
	r.POST("/kategori", controllers.CreateCategory)
	r.PUT("/kategori/:id", controllers.UpdateCategory)
	r.DELETE("/kategori/:id", controllers.DeleteCategory)

	// BOOK ROUTES
	r.GET("/buku", controllers.GetBooks)
	r.POST("/buku", controllers.CreateBook)
	r.PUT("/buku/:id", controllers.UpdateBook)
	r.DELETE("/buku/:id", controllers.DeleteBook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
