package user

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement;column:id_user"`
	TenantID     uint       `json:"tenant_id" gorm:"not null;column:tenant_id"`
	Username     string     `json:"username" gorm:"type:varchar(50);not null;column:username"`
	Picture      *string    `json:"picture,omitempty" gorm:"type:text;column:picture"`
	Email        string     `json:"email" gorm:"type:varchar(150);not null;column:email"`
	PasswordHash string     `json:"-" gorm:"type:varchar(150);not null;column:password"` // bcrypt hash
	Phone        string     `json:"phone" gorm:"type:varchar(30);column:phone"`
	FirstName    string     `json:"first_name" gorm:"type:varchar(100);column:first_name"`
	LastName     string     `json:"last_name" gorm:"type:varchar(100);column:last_name"`
	AddressID    uint       `json:"address_id" gorm:"column:address_id"`
	RoleID       uint       `json:"role_id" gorm:"not null;column:role_id"`
	TwoFAEnabled bool       `json:"two_fa_enabled" gorm:"default:false;column:two_fa_enabled"`
	Status       bool       `json:"status" gorm:"default:true;column:status"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    *time.Time `json:"update_at,omitempty" gorm:"autoUpdateTime;column:update_at"`
}

func (User) TableName() string { return "users" }
