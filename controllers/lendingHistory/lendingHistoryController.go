package lendingHistoryController

import (
	"fmt"
	"librarysystem/database"
	bookRepo "librarysystem/repository/book"
	lendingHistoryRepo "librarysystem/repository/lendingHistory"
	"librarysystem/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllLendingHistory(c *gin.Context) {
	var (
		result gin.H
	)

	students, err := lendingHistoryRepo.GetAllLendingHistory(database.DbConnection)
	fmt.Println(students)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": students,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertLendingHistory(c *gin.Context) {
	var lendHist structs.LendingHistory
	var book structs.Book

	err := c.ShouldBindJSON(&lendHist)
	if err != nil {
		panic(err)
	}

	// Validate Availbillity Book
	book.ID = lendHist.Book_id
	book.IsAvailable = true
	book, err = bookRepo.GetBookByID(database.DbConnection, book)
	if err != nil {
		panic(err)
	}
	if (book == structs.Book{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Buku tidak ditemukan atau sedang dipinjamkan",
		})
		return
	}

	// Update availbility of BOOK
	book.IsAvailable = false
	err = bookRepo.UpdateStatusBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	timeNow := time.Now().Format(time.RFC3339)
	lendHist.Created_at = timeNow
	err = lendingHistoryRepo.InsertLendingHistory(database.DbConnection, lendHist)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan data peminjaman",
	})
}

func DeleteLendingHistory(c *gin.Context) {

	var lendingHistory structs.LendingHistory
	id, _ := strconv.Atoi(c.Param("id"))

	lendingHistory.ID = int64(id)
	err := lendingHistoryRepo.DeleteLendingHistory(database.DbConnection, lendingHistory)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus LendingHistory",
	})
}
