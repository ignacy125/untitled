package main

import (
	"net/http"

	"fmt"
	"time"
	"strings"
	"html/template"
	"io/ioutil"
)

var login string
type Page struct{
	Title string
	Body []byte
}
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
//TODO tutaj też zmien na response i request
func witam(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("Witam na mojej stronie", request.URL.Path[1:])
	http.ServeFile(response, request, "logowanie.html")

}

func logowanie(response http.ResponseWriter, request *http.Request) {
	if (request.Method == "POST") {
		var correctPassword string
		request.ParseForm()

		fmt.Println("login", request.Form["login"])
		//fmt.Println("pass", request.Form["pass"])

		correctPassword = "12345"
		correctLogin := "aziron"

		login = request.Form["login"][0]
		haslo := request.Form["pass"][0]

		if strings.EqualFold(login, correctLogin) {
			if haslo == correctPassword {
				http.Redirect(response, request, "/internal", 302)

			} else {
				fmt.Println("Niepoprawne hasło")
				time.Sleep(2)
				http.Redirect(response, request, "/invalid_login", 302)
			}

		} else {
			fmt.Println("Niepoprawne login")
			time.Sleep(2)
			http.Redirect(response, request, "/invalid_login", 302)

		}


	} else{
		request.ParseForm()
		fmt.Println("Witam na mojej stronie 2", request.URL.Path)
	}
	//fmt.Fprintf(wysylacz,"Witaj")
	http.ServeFile(response, request, "logowanie.html")

}
func internalHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie 3")
	//http.ServeFile(wysylacz, pytanie, "internal")
	//p := "aziron"
	title := "aziron"
	//body :=[]byte{0}
	//p := loadPage("logowanie")
	p:= Page{Title: title}
	t := template.Must(template.ParseFiles("internal.html", "logowanie.html"))
	t.ExecuteTemplate(response, "internal", p)
	//t.Execute(response, p)
	//fmt.Fprintf(response, "Witaj %s", login)
}
func invalidLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "invalid_login.html")

}


func main() {
	http.HandleFunc("/", witam)

	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request){})
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/internal", internalHandler)
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}
}
