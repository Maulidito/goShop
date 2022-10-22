package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TotalPrice float64 `form:"total_price"`
	Cart       []Cart  `gorm:"foreignKey:Transaction_Fk"`
}
