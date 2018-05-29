package routers

import (
	"net/http"

	"../common"
	"github.com/gin-gonic/gin"
)

type Question struct {
	ID          int
	Description string
}

func TutorQuestionHistory(c *gin.Context) {

	// invoke db conn
	db, err := common.InitDB()
	if err != nil {
		return
	}
	defer db.Close()

	var (
		questions Question
		result    gin.H
	)

	id := c.Query("tutor_id")

	row := db.Where("answer_tutor_id = ?", id).Find(&questions)
	result = gin.H{
		"result": row,
	}
	c.JSON(http.StatusOK, result)
}
