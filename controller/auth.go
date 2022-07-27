package controller

import (
	"html/template"
	"log"
	"net/http"
	"simpleapp/model"

	"github.com/gorilla/securecookie"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseGlob("./view/index.gohtml")
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
		tmp.ExecuteTemplate(w, "index.gohtml", nil)
	}
}

// login handler

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")
		redirectTarget := "/"
		if email != "" && password != "" {
			// .. check credentials ..
			hash := model.SelectHash(email)
			match := CheckPasswordHash(password, hash)
			if match {
				setSession(email, w)
				redirectTarget = "/internal"
			}

		}
		http.Redirect(w, r, redirectTarget, http.StatusFound)
	}
}

// logout handler

func logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clearSession(w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// internal page

func internal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseGlob("./view/*.gohtml")
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}

		userName := getUserName(r)
		if userName != "" {
			tmp.ExecuteTemplate(w, "welcome.gohtml", nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}
