package user

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type createUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type updateUserRequest struct {
	Username string `json:"username" validate:"isdefault"`
	Password string `json:"password" validate:"required"`
}
