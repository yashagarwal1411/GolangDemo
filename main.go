package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/sessions"
	"gopkg.in/mailgun/mailgun-go.v1"
	"yash_agarwal/controllers"
	"yash_agarwal/conf"
	"yash_agarwal/apputils"
)

func main() {
	var err error
	conf.Db, err = sql.Open("mysql", conf.DSN)
	apputils.CheckErr(err)
	conf.MG = mailgun.NewMailgun(conf.MailGunDomain, conf.MailGunApiKey, conf.MailGunPubKey)
	conf.EncryptionKey = []byte(conf.EncryptionKeyString)
	conf.Store = sessions.NewCookieStore([]byte(conf.CookieStoreKey))

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.AuthHandler)
	r.HandleFunc("/profile", controllers.ProfileHandler)
	r.HandleFunc("/forgotPassword", controllers.ForgotPasswordHandler)
	r.HandleFunc("/resetPassword", controllers.ResetPasswordHandler)

	r.HandleFunc("/api/signin", controllers.SignInApi)
	r.HandleFunc("/api/googlesignin", controllers.GoogleSignIn)
	r.HandleFunc("/api/signup", controllers.SignUpApi)
	r.HandleFunc("/api/forgotPassword", controllers.SendForgotPasswordLink)
	r.HandleFunc("/api/resetPassword", controllers.ResetPassword)
	r.HandleFunc("/api/user/get", controllers.GetCurrentUser)
	r.HandleFunc("/api/user/update", controllers.UpdateCurrentUser)
	r.HandleFunc("/api/user/signout", controllers.SignOutApi)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./webapp/static/"))))
	http.Handle("/", r)

	fmt.Println("Server started and listening on " + conf.ServerPort)
	err = http.ListenAndServe(conf.ServerPort, r)
	apputils.CheckErr(err)
}