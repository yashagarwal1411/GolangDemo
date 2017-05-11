package dao

type ResetPasswordRequest struct {
	Email string
	Password string
	Token string
}
