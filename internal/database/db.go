package database

import (
	"fmt"
	"log"

	"github.com/b3nhard/car-rental/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(c config.DbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
