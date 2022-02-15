package handlers

import (
	"Ys7/config"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 请求错误
	if username != config.Username || password != config.Password {
		ErrorHandler(w, "用户名或密码错误", "401")
		return
	}

	// 请求成功
	log.Println("Login Success ", r.RequestURI)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: config.Username, Path: "/", Expires: time.Now().AddDate(20, 0, 0), HttpOnly: false})
	w.Write([]byte("ok"))
}
