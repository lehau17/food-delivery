package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSqlInstance() *gorm.DB {
	dsn := os.Getenv("MYSQL_CONNECT_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("-----CONNECTED MYSQL DATABASE------------")
	return db
}
