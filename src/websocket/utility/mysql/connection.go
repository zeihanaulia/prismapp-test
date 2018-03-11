package mysql

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func Connection(
	connectionString string,
) *gorm.DB {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	return db
}
