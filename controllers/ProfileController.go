package controllers

import (
	"net/http"
	"fmt"
	"yash_agarwal/conf"
	"yash_agarwal/apputils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, err := conf.Store.Get(r, "user-session")
	apputils.CheckErr(err)
	userId, _ := session.Values["userId"].(int64)
	if userId == 0 {
		w.WriteHeader(400)
		fmt.Print("ERROR IN FETCHING USER")
		w.Write([]byte("Please login first"))
		return
	}
	http.ServeFile(w, r, "webapp/profile.html")
}
