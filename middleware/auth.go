package auth

import (
	"fmt"
	"librarysystem/database"
	adminRepo "librarysystem/repository/admin"
	"librarysystem/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, _ := c.Request.BasicAuth()

	var userLog structs.User
	userLog.Username = user
	userLog.Password = password

	admin, err := adminRepo.GetAdminByUserPass(database.DbConnection, userLog)
	if err != nil {
		panic(err)
	}

	if (admin != structs.Admin{}) {
		c.Next()
	} else {
		fmt.Println("GAGAL")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
}
