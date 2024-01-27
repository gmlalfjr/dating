package models

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	FullName       string `json:"size:40;full_name"`
	Address        string `json:"size:256;address"`
	Age            int32  `json:"age"`
	Sex            string `json:"size:10;column:sex"`
}

type RegisterResponse struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
