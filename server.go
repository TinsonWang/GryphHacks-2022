package main

import "fmt"
import "io/ioutil"
import "log"
import "net/http"

func loginHandler(w http.ResponseWriter, r *http.Request) { 
}

func regHandler(w http.ResponseWriter, r *http.Request) {
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=Example");
    if err != nil {
        log.Fatalln(err);
    }

    body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Fatalln(err);
    }

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