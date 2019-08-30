package model

import "time"

// User 用户
type User struct {
	Id          string     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	ActivatedAt *time.Time `gorm:"column:activatedAt" json:"activatedAt"`
	Username    string     `gorm:"column:username" json:"username"`
	PassHash    string     `gorm:"column:passhash" json:"passHash"`
	Name        string     `gorm:"column:name" json:"name"`
	Email       string     `gorm:"column:email" json:"email"`
	Phone       string     `gorm:"column:phone" json:"phone"`
	Status      int        `gorm:"column:status" json:"status"`
}

func (u User) TableName() string {
	return "tb_main_user"
}
