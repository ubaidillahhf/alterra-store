package users

type UsersCreate struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersEdit struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	NewPassword     string `json:"newPassword"`
	RetypePassword  string `json:"retypePassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
