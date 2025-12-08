package user

import (
	"errors"

	"UserService/internal/common"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(req *CreateUserRequest) (*User, error)
	Update(id uint, req *UpdateUserRequest) (*User, error)
	Delete(id uint) error
	FindByTenant(tenantID uint) ([]User, error)
	ChangePassword(id uint, req *ChangePasswordRequest) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

// ------------------------- CREATE -------------------------
func (s *service) Create(req *CreateUserRequest) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(common.ERR_INTERNAL_ERROR)
	}

	u := &User{
		TenantID:     req.TenantID,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: string(hash),
		RoleID:       req.RoleID,
	}

	if err := s.repo.Create(u); err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}

	return u, nil
}

// ------------------------- UPDATE -------------------------
func (s *service) Update(id uint, req *UpdateUserRequest) (*User, error) {

	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New(common.ERR_NOT_FOUND)
	}

	u.Phone = req.Phone
	u.RoleID = req.RoleID

	if err := s.repo.Update(u); err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}

	return u, nil
}

// ------------------------- DELETE -------------------------
func (s *service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.New(common.ERR_DATABASE_ERROR)
	}
	return nil
}

// ------------------------- FIND BY TENANT -------------------------
func (s *service) FindByTenant(t uint) ([]User, error) {
	list, err := s.repo.FindByTenant(t)
	if err != nil {
		return nil, errors.New(common.ERR_DATABASE_ERROR)
	}
	return list, nil
}

// ------------------------- CHANGE PASSWORD -------------------------
func (s *service) ChangePassword(id uint, req *ChangePasswordRequest) error {

	u, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New(common.ERR_NOT_FOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.OldPassword)) != nil {
		return errors.New(common.ERR_INVALID_LOGIN)
	}

	newHash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	u.PasswordHash = string(newHash)

	return s.repo.Update(u)
}
