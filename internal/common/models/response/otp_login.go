package response

type LoginViaOtpVerificationResponse struct {
	UserId    string `json:"user_id"`
	AuthToken string `json:"auth_token"`
}
