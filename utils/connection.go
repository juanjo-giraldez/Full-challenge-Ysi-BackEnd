package utils

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

	func GetConnection() *gorm.DB { 

		db, err := gorm.Open("mysql", "user:admin@tcp(127.0.0.1:3306)/db_mysql?charset=utf8&parseTime=True&loc=Local")

		fmt.Println("db funcionando")

		if err != nil {
			log.Panic(err)
		}

		return db
	}