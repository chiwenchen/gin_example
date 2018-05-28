package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@unix(/tmp/mysql.sock)/snapask_development")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
