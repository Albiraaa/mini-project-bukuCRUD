package controllers

import (
	"net/http"
	"projek/database"
	"projek/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var body models.Category
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO categories(name) VALUES($1)", body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var body models.Category
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("UPDATE categories SET name=$1 WHERE id=$2", body.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
