package main

import "fmt"
// import "io/ioutil"
import "log"
import "net/http"
import "errors"
import "io"
import "os"

func loginHandler(w http.ResponseWriter, r *http.Request) { 
}

func regHandler(w http.ResponseWriter, r *http.Request) {
}

func downloadFile(URL, fileName string) error {
  //Get the response bytes from the url
  response, err := http.Get(URL)
  if err != nil {
    return err
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    return errors.New("Received non 200 response code")
  }
  //Create a empty file
  file, err := os.Create(fileName)
  if err != nil {
    return err
  }
  defer file.Close()

  //Write the bytes to the fiel
  _, err = io.Copy(file, response.Body)
  if err != nil {
    return err
  }

  return nil
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
    url := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=Example"
    resp, err := http.Get(url);
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("%+v\n", resp)

    fileName := "test.png"
    file := downloadFile(url, fileName)
    if file != nil {
      log.Fatal(file)
    }
    fmt.Printf("File %s download in current working directory", fileName)
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