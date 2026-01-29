package user

import (
	"errors"

	"UserService/internal/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Create(req *CreateUserRequest) (*User, error)
	Update(id uint, req *UpdateUserRequest) (*User, error)
	Delete(id uint) error
	FindByTenant(tenantID uint) ([]User, error)
	ChangePassword(id uint, req *ChangePasswordRequest) error
	CreateInternal(u *User) error
	HardDelete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(req *CreateUserRequest) (*User, error) {
	if req.TenantID == 0 || req.Email == "" || req.Password == "" || req.RoleID == 0 || req.Username == "" {
		return nil, errors.New(common.ERR_REQUIRED_FIELD)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(common.ERR_INTERNAL_ERROR)
	}

	u := &User{
		TenantID:     req.TenantID,
		Username:     req.Username,
		Picture:      req.Picture,
		Email:        req.Email,
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		AddressID:    req.AddressID,
		RoleID:       req.RoleID,
		PasswordHash: string(hash),
		TwoFAEnabled: false,
		Status:       true,
	}

	if err := s.repo.Create(u); err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}

	return u, nil
}

func (s *service) Update(id uint, req *UpdateUserRequest) (*User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(common.ERR_NOT_FOUND)
		}
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}

	if req.Username != nil {
		u.Username = *req.Username
	}
	if req.Picture != nil {
		u.Picture = req.Picture
	}
	if req.Phone != nil {
		u.Phone = *req.Phone
	}
	if req.FirstName != nil {
		u.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		u.LastName = *req.LastName
	}
	if req.AddressID != nil {
		u.AddressID = *req.AddressID
	}
	if req.RoleID != nil {
		u.RoleID = *req.RoleID
	}

	if err := s.repo.Update(u); err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}

	return u, nil
}

func (s *service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.New(common.ERR_DATABASE_ERROR)
	}
	return nil
}

func (s *service) FindByTenant(t uint) ([]User, error) {
	list, err := s.repo.FindByTenant(t)
	if err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}
	return list, nil
}

func (s *service) ChangePassword(id uint, req *ChangePasswordRequest) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New(common.ERR_NOT_FOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.OldPassword)) != nil {
		return errors.New(common.ERR_INVALID_LOGIN)
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(common.ERR_INTERNAL_ERROR)
	}

	u.PasswordHash = string(newHash)
	return s.repo.Update(u)
}

func (s *service) CreateInternal(u *User) error {
	u.Status = true
	return s.repo.Create(u)
}

func (s *service) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}
