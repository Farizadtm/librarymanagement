package book

import (
	"librarysystem/database"
	bookRepo "librarysystem/repository/book"
	"librarysystem/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := bookRepo.GetAllBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Format input data tidak sesuai",
		})
	}

	book.IsAvailable = true
	err = bookRepo.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Book Success",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Format input data tidak sesuai",
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = int64(id)
	err = bookRepo.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Update Book Success",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = int64(id)
	err := bookRepo.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete Book Success",
	})
}
