package types

type SignupPayload struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required"`
	Country   string `json:"country" validate:"required"`
}

type SigninPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	PINCODE  string `json:"pin_code"`
}
