package types

type UserPayload struct {
	Username  string `json:"username" validate:"required"`
	PINCODE   string `json:"pin_code" validate:"required"`
	IPAddress string `json:"ip_address" validate:"required"`
	Country   string `json:"country" validate:"required"`
}
