package model

import "time"

// User 用户
type User struct {
	Id           string     `gorm:"column:id" json:"id"`
	CreatedAt    time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	ActivatedAt  *time.Time `gorm:"column:activatedAt" json:"activatedAt"`
	Username     string     `gorm:"column:username" json:"username"`
	PassHash     string     `gorm:"column:passhash" json:"passHash"`
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	Phone        string     `gorm:"column:phone" json:"phone"`
	Avatar       string     `gorm:"column:avatar" json:"avatar"`
	Introduction string     `gorm:"column:introduction" json:"introduction"`
	Status       int        `gorm:"column:status" json:"status"`
}

func (u User) TableName() string {
	return "tb_main_user"
}

// secret
type Secret struct {
	Id        string    `gorm:"column:id" json:"id"`
	UserId    string    `gorm:"column:userId" json:"userId"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Name      string    `gorm:"column:name" json:"name"`
	Secret    string    `gorm:"column:secret" json:"secret"`
	Status    int       `gorm:"column:status" json:"status"`
}

func (u Secret) TableName() string {
	return "tb_main_secret"
}

// Role 角色实体
type Role struct {
	Id        string    `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name;size:100;index;"` // 角色名称
	Comment   string    `gorm:"column:comment;size:200;"`    // 备注
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}

func (r Role) TableName() string {
	return "tb_main_role"
}

// UserRole 用户角色关联实体
type UserRole struct {
	Id        string    `gorm:"column:id" json:"id"`
	UserID    string    `gorm:"column:userId;size:36;index;"` // 用户id
	RoleID    string    `gorm:"column:roleId;size:36;index;"` // 角色id
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}

func (r UserRole) TableName() string {
	return "tb_main_user_role"
}
