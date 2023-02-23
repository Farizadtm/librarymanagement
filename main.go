package main

import (
	// "github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"librarysystem/database"
	"os"

	adminControler "librarysystem/controllers/admin"
	bookController "librarysystem/controllers/book"
	lendingHistoryController "librarysystem/controllers/lendingHistory"
	returnHistoryController "librarysystem/controllers/returnHistory"
	studentController "librarysystem/controllers/student"
	auth "librarysystem/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("FAILED load file enviroment")
	} else {
		fmt.Println("Success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()

	// Student
	router.GET("/student", studentController.GetAllStudents)
	router.POST("/student", auth.BasicAuth, studentController.InsertStudent)
	router.PUT("/student/:id", auth.BasicAuth, studentController.UpdateStudent)
	router.DELETE("/student/:id", auth.BasicAuth, studentController.DeleteStudent)

	// Admin
	router.GET("/admin", auth.BasicAuth, adminControler.GetAllAdmin)
	router.POST("/admin", adminControler.InsertAdmin)
	router.PUT("/admin/:id", auth.BasicAuth, adminControler.UpdateAdmin)
	router.DELETE("/admin/:id", auth.BasicAuth, adminControler.DeleteAdmin)

	// Book
	router.GET("/book", bookController.GetAllBooks)
	router.POST("/book", auth.BasicAuth, bookController.InsertBook)
	router.PUT("/book/:id", auth.BasicAuth, bookController.UpdateBook)
	router.DELETE("/book/:id", auth.BasicAuth, bookController.DeleteBook)

	// History
	router.GET("/history", auth.BasicAuth, lendingHistoryController.GetAllLendingHistory)
	router.POST("/history", auth.BasicAuth, lendingHistoryController.InsertLendingHistory)
	router.DELETE("/history/:id", auth.BasicAuth, lendingHistoryController.DeleteLendingHistory)

	// Return History
	router.GET("/returnhistory", auth.BasicAuth, returnHistoryController.GetAllReturnHistory)
	router.POST("/returnhistory", auth.BasicAuth, returnHistoryController.InsertReturnHistory)

	router.Run(":" + os.Getenv("PORT"))
}
