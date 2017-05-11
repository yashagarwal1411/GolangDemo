package controllers

import (
	"net/http"
)

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webapp/resetpassword.html")
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webapp/forgotpassword.html")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webapp/auth.html")
}


