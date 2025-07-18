package auth

type CreateUserPayload struct {
	Username string `json:"username" validate:"required"`
}
