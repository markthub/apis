package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetDatabaseClient will return the database client instance of MySQL
func GetDatabaseClient(user, passw, addr, port, dbName string) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, passw, addr, port, dbName)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)

	return db
}
