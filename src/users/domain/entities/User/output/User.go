package output

type UserLoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	RoleName  string `json:"role_name"`
	Token     string `json:"token"`
}

type UserLoginDBResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleName  string `json:"role_name"`
	Token     string `json:"token"`
}
