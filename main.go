package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 error found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello....!!!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", hellohandler)

	log.Println("Starting server on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}