package models

import (
	"github.com/jinzhu/gorm"
	
)

type Person struct {
    gorm.Model
    Username string
    Email    string
	Address	 string
}
func (p *Person) TableName() string {
    return "persons" // Set the table name explicitly
}