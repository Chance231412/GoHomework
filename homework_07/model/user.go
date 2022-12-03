package model

type User struct {
	Name string `gorm:"primaryKey" form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone string `form:"phone" json:"phone" binding:"required"`
}
