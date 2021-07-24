package auth_service_wrapper

type TokenResponse struct {
	Token string `json:"token"`
}

type GetTokenResponseContract struct {
	Status bool
	Error  interface{}
	Result TokenResponse `json:"result"`
	Token  string        `json:"token"`
}

type AuthValidateTokenRequestContract struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	Authorized bool   `json:"authorized"`
	RequestID  string `json:"requestID"`
}

type AuthValidateTokenResponseContract struct {
	Status bool
	Error  interface{}
	Result ValidateTokenResponse `json:"result"`
}
