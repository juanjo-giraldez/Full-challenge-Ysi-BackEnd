package models

import (

	"gorm.io/gorm"
	
	) 

type Company struct {

	gorm.Model
	Name 		string 		`json:"Name"`
	Direction 	string 		`json:"Direction"`
	Phone 		string 		`json:"Phone"`
	Web 		string 		`json:"Web"`
	
}