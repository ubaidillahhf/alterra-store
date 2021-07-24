package users

type UserCreate struct {
	Name     string `json:"name" validate:"required,min=2"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserEdit struct {
	Name            string `json:"name" validate:"required,min=2"`
	Address         string `json:"address" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	NewPassword     string `json:"newPassword" validate:"required"`
	RetypePassword  string `json:"retypePassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
