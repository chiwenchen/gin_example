package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@unix(/tmp/mysql.sock)/snapask_development")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	type User struct {
		ID       int
		Username string
		Email    string
	}

	router := gin.Default()

	router.GET("/users/:id", func(c *gin.Context) {
		var (
			user   User
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, username, email from users where id = ?;", id)
		err = row.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			fmt.Println(err)
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": user,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)

	})

	router.GET("")

	router.Run(":3030")
}
