package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint8
}

//type User struct {
//	ID        uint
//	Name      string
//	Age       uint8
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}
