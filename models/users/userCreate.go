package users

type UserCreate struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEdit struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	NewPassword     string `json:"newPassword"`
	RetypePassword  string `json:"retypePassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
