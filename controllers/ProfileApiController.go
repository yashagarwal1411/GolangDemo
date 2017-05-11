package controllers

import (
	"net/http"
	"fmt"
	"yash_agarwal/models"
	"encoding/json"
	"yash_agarwal/conf"
	"yash_agarwal/apputils"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	userId, _ := session.Values["userId"].(int64)
	fmt.Println("USER ID IS" , userId)
	if userId == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		return
	}
	var user models.UserInfo
	row := conf.Db.QueryRow("SELECT * from userinfo where id=?", userId)
	apputils.CheckErr(err)
	apputils.UpdateUserStruct(row, &user)
	json.NewEncoder(w).Encode(user)
}

func UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.UserInfo
	err := decoder.Decode(&user)
	apputils.CheckErr(err)
	defer r.Body.Close()
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	user.Id, _ = session.Values["userId"].(int64)
	if user.Id == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		return
	}
	stmt, err := conf.Db.Prepare("UPDATE userinfo SET email=?, full_name=?, address=?, telephone=? where id=?")
	apputils.CheckErr(err)
	_, err = stmt.Exec(user.Email, user.FullName, user.Address, user.Telephone, user.Id)
	apputils.CheckErr(err)
	json.NewEncoder(w).Encode(user)
}