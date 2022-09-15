package auth

// struct for credential from API body request
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
