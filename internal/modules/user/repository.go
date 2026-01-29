package user

import "gorm.io/gorm"

type Repository interface {
	Create(u *User) error
	Update(u *User) error
	FindByID(id uint) (*User, error)
	FindByTenant(tenantID uint) ([]User, error)
	Delete(id uint) error     // soft delete => status=false
	HardDelete(id uint) error // hard delete real
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{db}
}

func (r *userRepository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) Update(u *User) error {
	return r.db.Save(u).Error
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var u User
	err := r.db.Where("id_user = ? AND status = true", id).First(&u).Error
	return &u, err
}

func (r *userRepository) FindByTenant(t uint) ([]User, error) {
	var list []User
	err := r.db.Where("tenant_id = ? AND status = true", t).Find(&list).Error
	return list, err
}

// ✅ soft delete
func (r *userRepository) Delete(id uint) error {
	return r.db.Model(&User{}).
		Where("id_user = ? AND status = true", id).
		Update("status", false).Error
}

func (r *userRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Where("id_user = ?", id).Delete(&User{}).Error
}
