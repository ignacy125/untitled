package main

import (
	"net/http"
	"fmt"
)


//TODO Stwórz nowy folder w katalogu projektu o nazwie /templates
//TODO Dobrze, ale pliki w folderze templates nie są dodane do systemu kontroli wersji "GIT"
//TODO aby dodać plik do GITa kliknij na niego prawym wybierz rozwijane menu Git -> a następnie "+ Add"
//TODO potem ctrl + k commit i push

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
	//TODO Czy te 2 uruchomienia są faktycznie takie same? Jeśli nie, to w funkcji main stwórz taki HandleFunc()
	//TODO żeby nie wyświetlać 2 razy komunikatu powitalnego na konsoli
	//TODO podpowiedź: przy pierwszym uruchomieniu ścieżka to "/" katakog główny
	//TODO przy drugim uruchomieniu ścieżka to "/favicon.ico" (przeglądarka internetowa szuka pliku ikony Twojej strony,
	//TODO która ma się pojawić na pasku zakładek lub w zapisanych (ulubionych) stronach.
	//TODO Komunikat z linii 24 wyświetla się na konsoli 2 razy, ponieważ funkcja "witam" uruchamia się przy podaniu
	//TODO ścieżki "/" oraz każdej innej nie obsłużonej przez HandleFunc() Jest to domyślna ścieżka.
	//TODO jeśli stworzysz HandleFunc który obsłuży ten przypadek i uruchomisz w nim funkcję która nic nie robi
	//TODO to komunikat nie będzie już wyświetlany 2 razy
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
		//TODO Jeśli masz problemy z tym przykładem, napisz maila z pytaniem, to pomogę Ci to rozwiązać
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

func HandleFuncHandler(response http.ResponseWriter, request *http.Request) {

}

func main() {
	//TODO do każdego szablonu stworzonego w folderze /template stwórz HandleFunc() oraz metode,
	//TODO która będzie ją obsługiwać tak jak "witam", "logowanie" "invalidLogin" itd... Niech każda metoda
	//TODO kończy się w nazwie słowem Handler
	http.HandleFunc("/", witam)
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/footer", footerHandler) //TODO skopiowałeś nazwę funkcji headerHandler i zapomniałeś pozmieniać nazwy :)
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/internal", internalHandler)
	http.HandleFunc("/logoutForm", logoutFormHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/", HandleFuncHandler)
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}

}



// Na zrobienie zadań masz czas do czwartku do 20:30. W razie pytań pisz maila,
// pamiętaj o stworzeniu brancha jeśli zamierzasz robić commit kodu, który się nie kompiluje