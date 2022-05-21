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

func insertDatabase(user_login string, user_password string) {
    host := goDotEnvVariable("HOST")
    database := goDotEnvVariable("DATABASE")
    user := goDotEnvVariable("DB_USER")
    password := goDotEnvVariable("DB_PASSWORD")

    // fmt.Printf("Host: %s\nDatabase: %s\nUser: %s\nPassword: %s\n", host, database, user, password)
    // fmt.Printf("LOGIN: %s\nPASSWORD: %s\n", user_login, user_password)

    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
    // fmt.Printf("CONNECTIONSTRING: %s\n", connectionString)

    // // Open up our database connection.
    // // I've set up a database on my local machine using phpmyadmin.
    // // The database is called testDb
    db, err := sql.Open("mysql", connectionString)

    // // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // // defer the close till after the main function has finished
    // // executing
    defer db.Close()

    // // perform a db.Query insert
    query := fmt.Sprintf(`INSERT INTO parkit (username, password, qrcode) VALUES ('%s', '%s', 1)`, user_login, user_password)
    insert, err := db.Query(query)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func flipQRCode(user_login string, user_password string) {
    host := goDotEnvVariable("HOST")
    database := goDotEnvVariable("DATABASE")
    user := goDotEnvVariable("DB_USER")
    password := goDotEnvVariable("DB_PASSWORD")
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    // // Open up database connection.
    db, err := sql.Open("mysql", connectionString)

    // // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // // defer the close till after the main function has finished
    defer db.Close()

    // // perform a db.Query insert
    query := fmt.Sprintf(`SELECT qrcode FROM parkit WHERE username = '%s'`, user_login)
    selected, err := db.Query(query)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }

    for selected.Next() {
        var name string
        if err := selected.Scan(&name); err != nil {
                log.Fatal(err)
        }
    }

    // // be careful deferring Queries if you are using transactions
    defer selected.Close()
}
