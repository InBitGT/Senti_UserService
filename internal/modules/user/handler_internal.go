package user

import (
	"UserService/internal/common"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type CreateAdminInternalRequest struct {
	TenantID  uint    `json:"tenant_id"`
	AddressID uint    `json:"address_id"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     string  `json:"phone"`
	RoleID    uint    `json:"role_id"`
	Picture   *string `json:"picture,omitempty"`
}

type CreateAdminInternalResponse struct {
	ID uint `json:"id"`
}

func (h *Handler) CreateAdminInternal(w http.ResponseWriter, r *http.Request) {
	var req CreateAdminInternalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if req.TenantID == 0 || req.Email == "" || req.Password == "" || req.RoleID == 0 || req.Username == "" {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	u := &User{
		TenantID:     req.TenantID,
		AddressID:    req.AddressID,
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hash),
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Picture:      req.Picture,
		RoleID:       req.RoleID,
		TwoFAEnabled: false,
		Status:       true,
	}

	if err := h.svc.CreateInternal(u); err != nil {
		common.ErrorResponse(w, http.StatusConflict, common.HTTP_CONFLICT, common.ERR_DUPLICATE, nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, CreateAdminInternalResponse{ID: u.ID}, common.HTTP_CREATED)
}
