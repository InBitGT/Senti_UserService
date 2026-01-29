package user

type CreateUserRequest struct {
	TenantID  uint    `json:"tenant_id"`
	Username  string  `json:"username"`
	Picture   *string `json:"picture,omitempty"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AddressID uint    `json:"address_id"`
	RoleID    uint    `json:"role_id"`
}

type UpdateUserRequest struct {
	Username  *string `json:"username,omitempty"`
	Picture   *string `json:"picture,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	AddressID *uint   `json:"address_id,omitempty"`
	RoleID    *uint   `json:"role_id,omitempty"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
