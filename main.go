package main

import (
	"fmt"
	"net/http"

	"./common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Username string
	Email    string
	LastName string
	Tos      bool
	SchoolId int
}

func main() {
	db, err := common.InitDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/users/:id", func(c *gin.Context) {
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

	})

	router.Run(":3030")
}
