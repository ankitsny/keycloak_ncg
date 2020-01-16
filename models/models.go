package models

// LoginModel :
type LoginModel struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// UserSignup :
type UserSignup struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lasttName,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
}
