package main

import (
	"net/http"

	"fmt"
)

func witam(wysylacz http.ResponseWriter, pytanie *http.Request) {
	fmt.Println("Witam na mojej stronie")
	http.ServeFile(wysylacz, pytanie, "logowanie.html")

}
func logowanie(wysylacz http.ResponseWriter, pytanie *http.Request) {
	if(pytanie.Method == "POST") {
	pytanie.ParseForm()
		fmt.Println("login", pytanie.Form["login"])
		fmt.Println("pass", pytanie.Form["pass"])

	}
	fmt.Println("Witam na mojej stronie")
	//fmt.Fprintf(wysylacz,"Witaj")
	http.ServeFile(wysylacz, pytanie, "logowanie.html")

}
func main(){
	http.HandleFunc("/", witam)
	http.HandleFunc("/logowanie", logowanie)
	error := http.ListenAndServe("localhost:5555", nil)
	if(error != nil){
		println(error)
	}
}
