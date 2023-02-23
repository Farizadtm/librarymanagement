package studentController

import (
	"librarysystem/database"
	studentRepo "librarysystem/repository/student"
	"librarysystem/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var (
		result gin.H
	)

	students, err := studentRepo.GetAllStudents(database.DbConnection)

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

func InsertStudent(c *gin.Context) {
	var student structs.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		panic(err)
	}

	err = studentRepo.InsertStudent(database.DbConnection, student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan Student",
	})
}

func UpdateStudent(c *gin.Context) {
	var student structs.Student
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&student)
	if err != nil {
		panic(err)
	}

	student.ID = int64(id)
	err = studentRepo.UpdateStudent(database.DbConnection, student)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Update Student",
	})

}

func DeleteStudent(c *gin.Context) {
	var student structs.Student
	id, _ := strconv.Atoi(c.Param("id"))

	student.ID = int64(id)
	err := studentRepo.DeleteStudent(database.DbConnection, student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus Student",
	})
}
