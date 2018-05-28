package main

import (
	"./routers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.GET("/users/:id", routers.FindUser)

	router.Run(":3030")
}
