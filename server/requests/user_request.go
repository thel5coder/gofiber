package requests

type UserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}
