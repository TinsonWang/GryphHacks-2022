package main

import "fmt"
import "log"
import "net/http"

func loginHandler(w http.ResponseWriter, r *http.Request) { 
}

func regHandler(w http.ResponseWriter, r *http.Request) {
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", regHandler)
	http.HandleFunc("/qr", qrHandler)
	http.HandleFunc("/map", mapHandler)

	fmt.Print("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}