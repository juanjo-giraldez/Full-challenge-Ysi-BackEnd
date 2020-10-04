package utils

import (
	"fmt"
	"../models"
)

func MigrateDB() {
	db := GetConnection()
	defer db.Close()

	fmt.Println("Migrating models....")
	db.AutoMigrate(&models.Company{})
}