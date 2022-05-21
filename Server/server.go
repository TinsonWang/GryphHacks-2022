package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "errors"
  "io"
  "os"
  "math/rand"
)

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
    min := 10000000000
    max := 30000000000
    rand := rand.Intn(max - min) + min
    url := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=" + string(rand)
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
    fileBytes, err := ioutil.ReadFile("test.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", regHandler)
	http.HandleFunc("/qr", qrHandler)
	http.HandleFunc("/map", mapHandler)

	fmt.Print("Starting server on port 5000\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
        log.Fatal(err)
    }
}