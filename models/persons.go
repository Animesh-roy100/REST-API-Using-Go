package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
