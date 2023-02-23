package adminControler

import (
	"fmt"
	"librarysystem/database"
	adminRepo "librarysystem/repository/admin"
	"librarysystem/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DateOnly = "2006-01-02"
)

func GetAllAdmin(c *gin.Context) {
	var (
		result gin.H
	)

	admins, err := adminRepo.GetAllAdmin(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": admins,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertAdmin(c *gin.Context) {
	var admin structs.Admin

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Format input data tidak sesuai",
		})
	}

	passLen := len([]rune(admin.Password))
	if passLen < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Password minimal 8 digit",
		})
		return
	}

	usernameLen := len([]rune(admin.Username))
	if usernameLen < 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Username minimal 4 digit",
		})
		return
	}
	_, err = time.Parse(DateOnly, admin.Birth_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Bad Request": "Input tanggal lahir tidak sesuai",
		})
		return
	}

	fmt.Println(admin)
	err = adminRepo.InsertAdmin(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Admin Success",
	})
}

func UpdateAdmin(c *gin.Context) {
	var admin structs.Admin

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		panic(err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	admin.ID = int64(id)
	err = adminRepo.UpdateAdmin(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Update Admin Success",
	})
}

func DeleteAdmin(c *gin.Context) {
	var admin structs.Admin
	id, _ := strconv.Atoi(c.Param("id"))

	admin.ID = int64(id)
	err := adminRepo.DeleteAdmin(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete Admin Success",
	})
}
