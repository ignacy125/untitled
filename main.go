package main

import (
"fmt"
"net/http"
"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		if k == "mail"{
			mail:=v[0]
			fmt.Println("mail:", mail)
		}
		if k == "pass"{
			pass:=v[0]
			fmt.Println("pass:", pass)
		}
		//fmt.Println("key:", k)
		//fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "<h1>iksd</h1>" +
		"<form action=\"index.html\" method=\"get\">" +
		"<input type=\"email\" name=\"mail\">" +
		"<input type=\"password\" name=\"pass\">" +
		"<input type=\"submit\" value=\"Submit\">" +
		"</form>") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}