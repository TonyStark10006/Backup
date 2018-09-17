package main

import (
	"net/http"
	"test/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))

	if user.Pasword == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		Cookie := http.Cookie{
			Name:     "_cookie",
			Vaule:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &Cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
