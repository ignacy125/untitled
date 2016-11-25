package main

import (
	"net/http"

	"fmt"
	"time"
	"strings"
)
var login string
//TODO tutaj też zmien na response i request
func witam(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie")
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


	}
	fmt.Println("Witam na mojej stronie")
	//fmt.Fprintf(wysylacz,"Witaj")
	http.ServeFile(response, request, "logowanie.html")

}
func internalHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie 2")
	//http.ServeFile(wysylacz, pytanie, "internal")
	fmt.Fprintf(response, "Witaj %s", login)
}
func invalidLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "invalid_login.html")

}


func main() {
	http.HandleFunc("/", witam)

	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/internal", internalHandler)
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}
}
