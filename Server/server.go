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
  "encoding/json"
)

type User struct {
  Login string
  Password string
}

// Handler for landing page
func mainHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Print("Welcome to ParkIt! Powered by FireTrucks.")
}

// This handler takes a login-password pair from the request object, finds it in the database, and flips the database column qrcode value
func loginHandler(w http.ResponseWriter, r *http.Request) {

  //First parse the request object
  var some_user User
  err := json.NewDecoder(r.Body).Decode(&some_user)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "POST request successful\n")
  fmt.Fprintf(w, "Login: %s\nPassword: %s\n", some_user.Login, some_user.Password)

  //Depending on login success or failure, prepare a response object
  rowCounts := flipQRCode(some_user.Login, some_user.Password)
  w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
  if (rowCounts == 0) {
    resp["Status"] = "Login Failed"
  } else {
    resp["Status"] = "Login Successful"
  }
	jsonResp, err2 := json.Marshal(resp)
	if err2 != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err2)
	}

  //Send response object
	w.Write(jsonResp)
	return
}

// This handler takes a login-password pair from the request object and adds it to the database
func regHandler(w http.ResponseWriter, r *http.Request) {

  //First parse the request object
  var some_user User
  err := json.NewDecoder(r.Body).Decode(&some_user)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "POST request successful\n")
  fmt.Fprintf(w, "Login: %s\nPassword: %s\n", some_user.Login, some_user.Password)

  //Call SQL function using data from request object
  insertDatabase(some_user.Login, some_user.Password)
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

  //Write the bytes to the file
  _, err = io.Copy(file, response.Body)
  if err != nil {
    return err
  }

  return nil
}

func qrHandler(w http.ResponseWriter, r *http.Request) {

  //Generate a random string for the URL
  min := 10000000000
  max := 30000000000
  rand := rand.Intn(max - min) + min
  url := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=" + string(rand)
  resp, err := http.Get(url);
  if err != nil {
      log.Fatalln(err)
  }
  fmt.Printf("%+v\n", resp)

  //Fetch the QR code image
  fileName := "test.png"
  file := downloadFile(url, fileName)
  if file != nil {
    log.Fatal(file)
  }
  fmt.Printf("File %s download in current working directory", fileName)

  //Upload the QR code image
  fileBytes, err := ioutil.ReadFile("test.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func main() {

  //Designate handlers for each endpoint
  http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", regHandler)
	http.HandleFunc("/qr", qrHandler)

  //Start server
	fmt.Print("Starting server on port 5000\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
        log.Fatal(err)
    }
}