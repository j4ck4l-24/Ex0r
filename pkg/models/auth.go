package models

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessfulLoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	UserId             int    `json:"user_id"`
	StoredHashPassword string `json:"hashed_password"`
	UserEmail          string `json:"email"`
	UserRole           string `json:"role"`
	DbUserName         string `json:"username"`
}