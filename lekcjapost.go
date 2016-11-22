package main

import (
	"net/http"

	"fmt"
)

func witam(wysylacz http.ResponseWriter, pytanie *http.Request) {
	fmt.Println("Witam na mojej stronie")
	http.ServeFile(wysylacz, pytanie, "logowanie.html")

}
//TODO zmien nazwy zmiennych wysylacz i pytanie na response i request
func logowanie(wysylacz http.ResponseWriter, pytanie *http.Request) {
	if(pytanie.Method == "POST") {
		//TODO Zapisz login i haslo do zmiennych np: login := pytanie.Form["login"] i nie wyswietlaj na konsoli hasla
		pytanie.ParseForm()
		fmt.Println("login", pytanie.Form["login"])
		fmt.Println("pass", pytanie.Form["pass"])
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
	http.ServeFile(wysylacz, pytanie, "logowanie.html")

}
func main(){
	http.HandleFunc("/", witam)
	//TODO podobnie jak /logowanie stworz nowa strone /internal a obsluge tej strony zapisz w funkcji internalHandler
	//TODO mozesz skorzystac z http.ServeFile tak jak w funkcji "witam"
	http.HandleFunc("/logowanie", logowanie)
	//TODO Zmien nazwe zmiennej error, poniewaz nazwa koliduje ze zmienna error wbudowana w jezyk golang
	error := http.ListenAndServe("localhost:5555", nil)
	if(error != nil){
		println(error)
	}
}


//TODO Jesli masz jakies pytanie albo czegos nie wiesz, napisz maila do mnie. Czas na wykonanie zadania do czwartku do 20:00
//TODO Po skonczonej pracy w danym dniu uzyj gita CTRL + k i wrzuc zmiany na githuba