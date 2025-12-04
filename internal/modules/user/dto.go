package user

type CreateUserRequest struct {
	TenantID uint   `json:"tenant_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type UpdateUserRequest struct {
	Phone  string `json:"phone"`
	RoleID uint   `json:"role_id"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
