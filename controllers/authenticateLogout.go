package controllers

import (
	"net/http"
	"time"
)

func ControllerAuthenticateLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie {
		Path: "/",
		Name: "sid",
		Expires: time.Now().Add(-100 * time.Hour),
		MaxAge: -1 }

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}