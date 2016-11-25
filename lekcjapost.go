package main

import (
	"net/http"

	"fmt"
)
//TODO tutaj też zmien na response i request
func witam(wysylacz http.ResponseWriter, pytanie *http.Request) {
	fmt.Println("Witam na mojej stronie")
	http.ServeFile(wysylacz, pytanie, "logowanie.html")

}
func logowanie(response http.ResponseWriter, request *http.Request) {
	if(request.Method == "POST") {
		//TODO Zapisz login i haslo do zmiennych np: login := pytanie.Form["login"] i nie wyswietlaj na konsoli hasla
		//TODO nie zapisałeś hasła i loginu do zmiennych, a hasło nadal jest wyświetlana na konsoli
		request.ParseForm()
		fmt.Println("login", request.Form["login"])
		fmt.Println("pass", request.Form["pass"])
		//TODO tak jak pisałem w smsie, brakuje operatora przypisania ":="
		var correctPassword = 12345
		var correctLogin = "aziron"

		if //TODO tego if'a możesz zakomentować jeśli nie wiesz jak go dokończyć
		//TODO Stworz zmienne correctPassword i correctLogin,
		//TODO Przypisz im wartosci ktore sobie wymyslisz aby za ich pomoca mozna bylo sie zalogowac na strone
		//TODO Za pomoca warunku if sprawdz czy podany w formularzu login i haslo pasuja do correctPassword itd...
		//TODO przyklad warunku if: https://gobyexample.com/if-else
		//TODO login moze byc caseInsensitive, czyli "aziron" to to samo co "Aziron" (wielkosc liter nie ma znaczenia)
		//TODO Do tego sluzy metoda EqualFold w klasie strings. Przykład użycia takiej metody jest ponizej
		//TODO http://stackoverflow.com/questions/30196780/case-insensitive-string-comparison-in-go
		//TODO Po poprawnym zalogowaniu przekieruj uzytkownika na strone /internal (pamietaj ze najpierw musisz ja utworzyc)
		//TODO Przekierowanie robi sie tak: http.Redirect(response, request, "/sciezka_do_strony_internal", 302)
		//TODO jesli haslo lub login sa niepoprawne przekieruj uzytkownika na strone /logowanie

	}
	fmt.Println("Witam na mojej stronie")
	//fmt.Fprintf(wysylacz,"Witaj")
	http.ServeFile(response, request, "logowanie.html")

}
//TODO tutaj też zmien na response i request
func internalHander(wysylacz http.ResponseWriter, pytanie *http.Request) {
	fmt.Println("Witam na mojej stronie 2")
	http.ServeFile(wysylacz, pytanie, "internal")




}
func main(){
	http.HandleFunc("/", witam)

	//TODO mozesz skorzystac z http.ServeFile tak jak w funkcji "witam"
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/internal") //TODO brakuje wywołania funkcji "internalHandler", która ma obsłużyć nową stronę, tak jak wyżej "logowanie"
	err := http.ListenAndServe("localhost:5555", nil)
	if(err != nil){
		//TODO Dobrze, ale println teraz nie wyswietli tego błędu, bo używasz starej nazwy zmiennej
		println(error)
	}
}


//TODO Jesli masz jakies pytanie albo czegos nie wiesz, napisz maila do mnie. Czas na wykonanie zadania do czwartku do 20:00
//TODO Po skonczonej pracy w danym dniu uzyj gita CTRL + k i wrzuc zmiany na githuba

//TODO ogólnie fajnie że dużo zrobiłeś, ale zostawiłeś kod w stanie w którym się nie kompiluje. Lepiej zakomentować te fragmenty :)