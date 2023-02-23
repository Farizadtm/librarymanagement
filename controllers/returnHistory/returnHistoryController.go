package returnHistoryController

import (
	"librarysystem/database"
	bookRepo "librarysystem/repository/book"
	returnHistoryRepo "librarysystem/repository/returnHistory"
	"librarysystem/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllReturnHistory(c *gin.Context) {
	var (
		result gin.H
	)

	returnHistory, err := returnHistoryRepo.GetAllReturnHistory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": returnHistory,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertReturnHistory(c *gin.Context) {
	var returnHist structs.ReturnHistory
	var book structs.Book

	err := c.ShouldBindJSON(&returnHist)
	if err != nil {
		panic(err)
	}

	// Validate Availbillity Book
	book.ID = returnHist.Book_id
	book.IsAvailable = true
	dataBook, err := bookRepo.GetBookByID(database.DbConnection, book)
	if err != nil {
		panic(err)
	}
	if (dataBook != structs.Book{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Buku tidak ditemukan atau sedang tidak dipinjamkan",
		})
		return
	}

	// Update availbility of BOOK
	err = bookRepo.UpdateStatusBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	timeNow := time.Now().Format(time.RFC3339)
	returnHist.Created_at = timeNow
	err = returnHistoryRepo.InsertReturnHistory(database.DbConnection, returnHist)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan data peminjaman",
	})
}
