package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var mytemp *template.Template

type userstuff struct {
	Name, Email, Message string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/formhandler", formhandler)
	http.HandleFunc("/link1", linkone)
	//http.HandleFunc("/link2", linktwo)
	//http.HandleFunc("/link3", linkthree)
	mytemp = template.Must(template.ParseFiles("form.html", "formhandler.html", "link1.html" /*,"link2.html","link3.html"*/))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := mytemp.ExecuteTemplate(w, "formhandler.html", struct {
		Name    string
		Email   string
		Message string
	}{r.FormValue("name"), r.FormValue("email"), r.FormValue("msg")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func linkone(w http.ResponseWriter, r *http.Request) {
	thisuser := r.FormValue("name") + " " + r.FormValue("email") + " " + r.FormValue("message")
	err := mytemp.ExecuteTemplate(w, "link1.html", thisuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
