package model

import "time"

// User 用户
type User struct {
	Id          string     `gorm:"column:id" json:"id"`
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

// secret
type Secret struct {
	Id          string     `gorm:"column:id" json:"id"`
	UserId          string     `gorm:"column:userId" json:"userId"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	Name        string     `gorm:"column:name" json:"name"`
	Secret      string     `gorm:"column:secret" json:"secret"`
	Status      int        `gorm:"column:status" json:"status"`
}

func (u Secret) TableName() string {
	return "tb_main_secret"
}