package user

type User struct {
	ID           uint   `gorm:"primaryKey"`
	TenantID     uint   `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Phone        string
	RoleID       uint `gorm:"not null"`
	TwoFAEnabled bool `gorm:"default:false"`
	IsActive     bool `gorm:"default:true"`
}
