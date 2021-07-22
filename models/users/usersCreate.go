package users

type UsersCreate struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
