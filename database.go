package main

import (
    "os"
    "fmt"
    "godotenv"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
      log.Fatalf("Error loading .env file")
    }

    return os.Getenv(key)
  }

func main() {
    host := goDotEnvVariable("HOST")
    database := goDotEnvVariable("DATABASE")
    user := goDotEnvVariable("DB_USER")
    password := goDotEnvVariable("DB_PASSWORD")

    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", connectionString)

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO parkit (username, password, qrcode) VALUES ( 'tinson-tester', 'tinson-tester-password', 1)")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}