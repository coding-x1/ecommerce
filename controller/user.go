package controller

import (
	"log"
	"net/http"
	"simpleapp/model"
	"text/template"
)

func Master() {
	//Application logic.
	http.HandleFunc("/", index())

	http.HandleFunc("/newUser", newUser())
	http.HandleFunc("/newUserProcess", newUserProcess())

	http.HandleFunc("/internal", internal())

	http.HandleFunc("/login", login())
	http.HandleFunc("/logout", logout())
}

func newUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseGlob("./view/newUser.gohtml")
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
		tmp.ExecuteTemplate(w, "newUser.gohtml", nil)
	}
}

func newUserProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fName := r.FormValue("fName")
		lName := r.FormValue("lName")
		email := r.FormValue("email")
		password := randSeq(10)
		hash, _ := HashPassword(password)
		user := []string{fName, lName, email, hash}
		model.CreateUser(user)
		mailDetails := map[string]string{
			"to":      email,
			"subject": fName + " You can login",
			"body":    "Your password to login is : " + password,
		}
		sendMail(mailDetails)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
