package common

import (
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@unix(/tmp/mysql.sock)/snapask_development")
	return db, err
}
