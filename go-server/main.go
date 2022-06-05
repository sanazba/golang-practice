package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	phone := r.FormValue("phone")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Phone = %s\n", phone)
	log.Println(name, address, phone)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bye" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	log.Println("handeling bye request")
	fmt.Fprintf(w, "bye")
}

func main() {
	fileserver := http.FileServer(http.Dir("./go-server/static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)

	fmt.Println("starting server at port 8085")
	err := http.ListenAndServe(":8085", nil)
	fmt.Println(err)
	if err != nil {
		fmt.Println("hi")
		log.Fatal(err)
	}
	// //	if err := http.ListenAndServe(":8085", nil); err != nil {
	// 		log.Fatal(err)
	//	}
}
