package router

import (
	"Ys7/config"
	"Ys7/internal/router/handlers"
	"log"
	"net/http"
)

func CheackCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil || cookie.Value != config.Username {
			log.Printf("Cookie check failed: %s\n", r.RequestURI)
			handlers.ErrorHandler(w, "身份验证失败", "401")
		} else {
			log.Println("Authorization Success ", r.RequestURI)
			next.ServeHTTP(w, r)
		}
	})
}

func SetCorsHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.106:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
