package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `form:"name"`
	Password string    `form:"password"`
	Email    string    `form:"email"`
	Cart     []Cart    `gorm:"foreignKey:User_Fk"`
	Product  []Product `gorm:"foreignKey:User_Fk"`
}
