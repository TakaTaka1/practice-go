package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else {

		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
		http.Redirect(w, r, "/welcome", http.StatusFound)
	}
}

// func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
// 	t, err := template.ParseFile(tmpl+".html", nil)
// 	if err != nil {
// 		http.Error(w, err.String(), http.StatusInternalServerError)
// 		return
// 	}
// 	err = t.Execute(w, p)
// 	if err != nil {
// 		http.Error(w, err.String(), http.StatusInternalServerError)
// 	}
// }

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./welcome.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9001",
	}
	//http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	server.ListenAndServe()
}
