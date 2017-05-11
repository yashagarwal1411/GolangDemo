package controllers

import (
	"net/http"
	"encoding/json"
	"yash_agarwal/models"
	"fmt"
	"yash_agarwal/dao"
	"yash_agarwal/apputils"
	"yash_agarwal/conf"
	"google.golang.org/api/oauth2/v1"
	"gopkg.in/mailgun/mailgun-go.v1"
)

func SignInApi(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var user models.UserInfo
	err := decoder.Decode(&user)
	fmt.Println(user.Email + " " + user.Password)
	apputils.CheckErr(err)
	defer r.Body.Close()
	row := conf.Db.QueryRow("SELECT id, email, password, address, telephone, full_name from userinfo where email=? and password=?", user.Email, user.Password)
	apputils.CheckErr(err)
	apputils.UpdateUserStruct(row, &user)
	if user.Id == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		return
	}
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	session.Values["userId"] = user.Id
	fmt.Println("User id stored in session is: ", user.Id)
	session.Save(r, w)
	json.NewEncoder(w).Encode(user)
}


func SignUpApi(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var user models.UserInfo
	err := decoder.Decode(&user)
	apputils.CheckErr(err)
	defer r.Body.Close()
	stmt, err := conf.Db.Prepare("INSERT into userinfo (email, password) VALUES (?, ?)")
	apputils.CheckErr(err)

	res, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	user.Id, err = res.LastInsertId()
	apputils.CheckErr(err)
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	session.Values["userId"] = user.Id
	fmt.Println("User id stored in session is: " + (string)(user.Id))
	session.Save(r, w)
	json.NewEncoder(w).Encode(user)
}

func GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var idToken  dao.GoogleSignInRequest
	err := decoder.Decode(&idToken)
	apputils.CheckErr(err)
	defer r.Body.Close()
	var httpClient = &http.Client{}
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := *oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken.IdToken)
	tokenInfo, err := tokenInfoCall.Do()
	apputils.CheckErr(err)
	userEmail := tokenInfo.Email
	if userEmail == "" {
		w.WriteHeader(400)
		fmt.Println("ERROR IN FETCHING USER")
		return
	}
	var user models.UserInfo
	user.Email = userEmail
	fmt.Println("Email for google sign in user is: " + userEmail)
	row := conf.Db.QueryRow("SELECT id, email, password, address, telephone, full_name from userinfo where email=?", user.Email)
	apputils.CheckErr(err)
	apputils.UpdateUserStruct(row, &user)
	if user.Id == 0 {
		fmt.Println("Creating new user for email: " + userEmail)
		user.Password = "erwrewrwhrekwjhrkhwrke" //Some random password
		stmt, err := conf.Db.Prepare("INSERT into userinfo (email, password) VALUES (?, ?)")
		apputils.CheckErr(err)

		res, err := stmt.Exec(user.Email, user.Password)
		apputils.CheckErr(err)

		user.Id, err = res.LastInsertId()
		apputils.CheckErr(err)
	}
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	session.Values["userId"] = user.Id
	fmt.Println("User id stored in session is: ", user.Id)
	session.Save(r, w)
	json.NewEncoder(w).Encode(user)
}


func SignOutApi(w http.ResponseWriter, r *http.Request) {
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	session.Values["userId"] = nil
	session.Save(r, w)
	var successResp dao.SucccessResp
	successResp.Message = "Successfully logged out"
	json.NewEncoder(w).Encode(successResp)
}

func SendForgotPasswordLink(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestData  dao.SendForgotPasswordLinkRequest
	err := decoder.Decode(&requestData)
	apputils.CheckErr(err)
	var user models.UserInfo
	row := conf.Db.QueryRow("SELECT id, email, password, address, telephone, full_name from userinfo where email=?", requestData.Email)
	apputils.CheckErr(err)
	apputils.UpdateUserStruct(row, &user)
	if user.Id == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		return
	}
	ciphertext, err := apputils.Encrypt([]byte(user.Password), conf.EncryptionKey)
	apputils.CheckErr(err)

	forgotPasswordLink := conf.ServerUrl + "/resetPassword?email="+ user.Email + "&token=" + ciphertext
	fmt.Println("Forgot pass link: " + forgotPasswordLink)

	forgotPasswordBody := `
	Hi,

	Click on the following link to reset your password:
	<a href="` + forgotPasswordLink + `">Click Here</a>

	Regards,
	Golang Assignment Team
	`
	message := mailgun.NewMessage(
		"sender@example.com",
		"Forgot Password",
		forgotPasswordBody,
		"testgolangassignment@mailinator.com")
	resp, id, err := conf.MG.Send(message)
	apputils.CheckErr(err)
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	var successResp dao.SucccessResp
	successResp.Message = "Successfully logged out"
	json.NewEncoder(w).Encode(successResp)
}



func ResetPassword(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestData  dao.ResetPasswordRequest
	err := decoder.Decode(&requestData)
	apputils.CheckErr(err)
	pass, err := apputils.Decrypt(requestData.Token, conf.EncryptionKey)
	apputils.CheckErr(err)
	var user models.UserInfo
	fmt.Println(requestData.Email + " " + requestData.Password)
	apputils.CheckErr(err)
	defer r.Body.Close()
	row := conf.Db.QueryRow("SELECT id, email, password, address, telephone, full_name from userinfo where email=? and password=?", requestData.Email, pass)
	apputils.CheckErr(err)
	apputils.UpdateUserStruct(row, &user)
	if user.Id == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		return
	}
	stmt, err := conf.Db.Prepare("UPDATE userinfo SET password=? where id=?")
	apputils.CheckErr(err)
	_, err = stmt.Exec(requestData.Password, user.Id)
	apputils.CheckErr(err)
	var successResp dao.SucccessResp
	successResp.Message = "Password changed Successfully"
	json.NewEncoder(w).Encode(successResp)
}