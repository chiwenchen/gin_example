package main

import (
	"./routers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.GET("/users/:id", routers.FindUser)
	router.GET("/tutors/question_history", routers.TutorQuestionHistory)

	router.Run(":3030")
}
