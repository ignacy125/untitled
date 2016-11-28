package main
//TODO Stwórz nowy folder w katalogu projektu o nazwie /templates, będą w nim trzymane szablony stron internetowych
//TODO w tym katalogu stwórz pliki html o nazwach "header", "footer", "loginForm", "logoutForm" "internal", "default", "welcome"
//TODO na lekcji będziemy korzystać z tych plików

import (
	"net/http"

	"fmt"
	"time"
	"strings"
)
var login string

func witam(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie", request.URL.Path)

	//TODO zauważyłeś że komunikat się wyświetla w konsoli 2 razy? Nie wiesz dlaczego tak jest?
	//TODO wyświetl na konsoli oprócz standardowego powitania wartość request.URL.Path.
	//TODO żeby to zrobić, musisz najpierw użyć request.ParseForm()
	//TODO Czy te 2 uruchomienia są faktycznie takie same? Jeśli nie, to w funkcji main stwórz taki HandleFunc()
	//TODO żeby nie wyświetlać 2 razy komunikatu powitalnego na konsoli
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
		//TODO popraw warunek sprawdzania poprawności loginu i hasła tak, aby wszystko mieściło się w
		//TODO jednej instrukcji if{} else{} W przykładzie poniżej drugi "if" pokazuje jak to zrobić"
		//TODO Przykład: https://www.quora.com/What-is-the-difference-between-using-and-in-golang-IF-statements
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

func headerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/header.html")

}

func footerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/footer.html")

}
/*func internalHandler(response http.ResponseWriter, request *http.Request) {
	//http.ServeFile(response, request, "templates/internal.html")

} */
func welcomeHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/welcome.html")

}
func defaultHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/default.html")

}
func logoutFormHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/logoutForm.html")

}


func main() {
	//TODO do każdego szablonu stworzonego w folderze /template stwórz HandleFunc() oraz metode,
	//TODO która będzie ją obsługiwać tak jak "witam", "logowanie" "invalidLogin" itd... Niech każda metoda
	//TODO kończy się w nazwie słowem Handler
	http.HandleFunc("/", witam)
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/footer", headerHandler)
	http.HandleFunc("/default", headerHandler)
	http.HandleFunc("/internal", headerHandler)
	http.HandleFunc("/logoutForm", headerHandler)
	http.HandleFunc("/welcome", headerHandler)
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}
}



// Na zrobienie zadań masz czas do czwartku do 20:30. W razie pytań pisz maila,
// pamiętaj o stworzeniu brancha jeśli zamierzasz robić commit kodu, który się nie kompiluje