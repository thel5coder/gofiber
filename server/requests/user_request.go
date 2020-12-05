package requests

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
