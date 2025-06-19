package model

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	Profil    *string   `gorm:"column:profil"`
}

func (User) TableName() string {
	return "user"
}
