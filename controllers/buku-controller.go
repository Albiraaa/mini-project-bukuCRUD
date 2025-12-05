package controllers

import (
	"net/http"
	"projek/database"
	"projek/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT books.id, books.title, books.author, categories.id, categories.name
		FROM books
		JOIN categories ON books.category_id = categories.id
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.CategoryID, &b.Category)
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var b models.Book

	err := database.DB.QueryRow(`
		SELECT books.id, books.title, books.author, categories.id, categories.name
		FROM books
		JOIN categories ON books.category_id = categories.id
		WHERE books.id = $1
	`, id).Scan(&b.ID, &b.Title, &b.Author, &b.CategoryID, &b.Category)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

func CreateBook(c *gin.Context) {
	var body models.Book
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec(
		"INSERT INTO books(title, author, category_id) VALUES($1, $2, $3)",
		body.Title, body.Author, body.CategoryID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var body models.Book
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec(
		"UPDATE books SET title=$1, author=$2, category_id=$3 WHERE id=$4",
		body.Title, body.Author, body.CategoryID, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
