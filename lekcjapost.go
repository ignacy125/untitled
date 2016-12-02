package main
import (
	"net/http"
	"fmt"
	"time"

	"strings"
	"html/template"

)
var login string


func witam(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie", request.URL.Path)

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


		if strings.EqualFold(login, correctLogin) && haslo == correctPassword {
			fmt.Println("Zalogowano")
			http.Redirect(response, request, "/internal", 302)
		} else {
			fmt.Println("Hasło lub login są niepoprawne")
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

func headerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/header.html")

}

func footerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/footer.html")

}

func welcomeHandler(response http.ResponseWriter, request *http.Request) {
	type Dane struct{
		Title string
		User string

	}
	data := Dane{Title:"Welcome"}
	//http.ServeFile(response, request, "templates/welcome.html")
	t := template.Must(template.ParseFiles("templates/welcome.html"))
	t.ExecuteTemplate(response, "welcome", data)

}
func defaultHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/default.html")

}
func logoutFormHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/logoutForm.html")

}



func main() {
	http.HandleFunc("/", witam)
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/footer", footerHandler)
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/internal", internalHandler)
	http.HandleFunc("/logoutForm", logoutFormHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/favicon.ico", func (response http.ResponseWriter, request *http.Request) {})
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}

}



// Na zrobienie zadań masz czas do czwartku do 20:30. W razie pytań pisz maila,
// pamiętaj o stworzeniu brancha jeśli zamierzasz robić commit kodu, który się nie kompiluje