package user

import "gorm.io/gorm"

type Repository interface {
	Create(u *User) error
	Update(u *User) error
	FindByID(id uint) (*User, error)
	FindByTenant(tenantID uint) ([]User, error)
	Delete(id uint) error
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
	err := r.db.First(&u, id).Error
	return &u, err
}

func (r *userRepository) FindByTenant(t uint) ([]User, error) {
	var list []User
	err := r.db.Where("tenant_id = ?", t).Find(&list).Error
	return list, err
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
