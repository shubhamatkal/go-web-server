package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not available", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello !")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submitted Successfully")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("server started at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
