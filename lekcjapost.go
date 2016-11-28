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
var correctLogin string = "aziron"
type User struct{
	Name string
	Login string
}

type Page struct{
	Title string
	Body []byte
	User User
}
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
//TODO tutaj też zmien na response i request
func witam(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	path:= request.URL.Path
	if path != "/" {
		defaultHandler(response, request)
	} else {
		fmt.Println("Witam na mojej stronie", path)
		page:= Page{Title: "welcome"}
		getTemplate(response, request, page, "welcome")
	}
}

func logowanie(response http.ResponseWriter, request *http.Request) {
	if (request.Method == "POST") {
		var correctPassword string
		request.ParseForm()

		fmt.Println("login", request.Form["login"])
		//fmt.Println("pass", request.Form["pass"])

		correctPassword = "12345"

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
	title := "internal"
	if login == correctLogin {
		user:= User{"Ernest", login}
		page:= Page{Title: title, User: user}
		getTemplate(response, request, page, "internal")
	} else{
		page:= Page{Title: title}
		getTemplate(response, request, page, "welcome")
	}
}
func invalidLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "invalid_login.html")
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	login = ""
	internalHandler(response, request)
}

func defaultHandler(response http.ResponseWriter, request *http.Request) {
	title := request.URL.Path
	page:= Page{Title: title}
	getTemplate(response, request, page, "default")
}

func getTemplate(response http.ResponseWriter, request *http.Request, page Page, tmpl string){
	request.ParseForm()
	t := template.Must(template.ParseGlob("templates/*"))
	t.ExecuteTemplate(response, tmpl, page)
}


func main() {
	http.HandleFunc("/", witam)
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/favicon.ico", func (w http.ResponseWriter, r *http.Request){})
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/internal", internalHandler)
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}
}
