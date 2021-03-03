package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"column:name;size:255" binding:"required" json:"Name"`
	Occupation     string `gorm:"column:occupation;size:255" binding:"required" json:"Occupation"`
	Email          string `gorm:"column:email;size:110;unique" binding:"required,email" json:"Email"`
	Password       string `gorm:"column:password" binding:"required" json:"Password"`
	//AvatarFileName string `gorm:"default:'user.png'" json:"Avatar_file_name"`
	AvatarFileName string `gorm:"column:avatar_file_name" json:"Avatar_file_name"`
	Role           string `gorm:"column:role;default:'user'" json:"Role"`
}

type UserFormatter struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token 	   string `json:"token"`
}

// Custom Name table
//func (e User) TableName() string {
//	return "user"
//}
