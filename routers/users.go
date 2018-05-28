package routers

import (
	"net/http"

	"../common"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int
	Username string
	Email    string
	LastName string
	Tos      bool
	SchoolId int
}

func FindUser(c *gin.Context) {
	db, err := common.InitDB()
	if err != nil {
		return
	}
	defer db.Close()

	var (
		users  User
		result gin.H
	)
	id := c.Param("id")
	row := db.First(&users, id)
	result = gin.H{
		"result": row,
	}
	c.JSON(http.StatusOK, result)
}
