package main

import (
    "os"
    "fmt"
    "godotenv"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// use godot package to load/read the .env file and return value of key
func goDotEnvVariable(key string) string {

    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
      log.Fatalf("Error loading .env file")
    }

    return os.Getenv(key)
  }

//Registers a new user using the provided login and password
func insertDatabase(user_login string, user_password string) {

    //Load in variables
    host := goDotEnvVariable("HOST")
    database := goDotEnvVariable("DATABASE")
    user := goDotEnvVariable("DB_USER")
    password := goDotEnvVariable("DB_PASSWORD")

    // fmt.Printf("Host: %s\nDatabase: %s\nUser: %s\nPassword: %s\n", host, database, user, password)
    // fmt.Printf("LOGIN: %s\nPASSWORD: %s\n", user_login, user_password)

    //Set up string variable
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
    // fmt.Printf("CONNECTIONSTRING: %s\n", connectionString)

    //Establish database connection.
    db, err := sql.Open("mysql", connectionString)

    //if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    //Close connection after finishing
    defer db.Close()

    //Prepare a query and perform an INSERT query
    query := fmt.Sprintf(`INSERT INTO parkit (username, password, qrcode) VALUES ('%s', '%s', 1)`, user_login, user_password)
    insert, err := db.Query(query)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }

    //Close after finishing
    defer insert.Close()
}

//Upon a valid login, flips the QR code status in the database, indicating active or inactive QR code
func flipQRCode(user_login string, user_password string)(rowCounts int) {

    //Load in variables
    host := goDotEnvVariable("HOST")
    database := goDotEnvVariable("DATABASE")
    user := goDotEnvVariable("DB_USER")
    password := goDotEnvVariable("DB_PASSWORD")

    //Set up string variable
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    //Establish database connection.
    db, err := sql.Open("mysql", connectionString)

    //if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    //Close connection after finishing
    defer db.Close()

    //Prepare and perform a SELECT query
    query := fmt.Sprintf(`SELECT qrcode FROM parkit WHERE username='%s' AND password='%s'`, user_login, user_password)
    fmt.Printf(query)
    rows, err := db.Query(query)

    //if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }

    //Close after finishing
    defer rows.Close()

    //Tracker to determine if login success or failure
    rowCounts = 0
    for rows.Next() {
        rowCounts++
        var qrcode int
        if err := rows.Scan(&qrcode); err != nil {
                log.Fatal(err)
        }
        //Flip qrcode status based on its current value
        if (qrcode == 1) {

            //Prepare and perform an UPDATE query
            query2 := fmt.Sprintf(`UPDATE parkit SET qrcode=0 WHERE username='%s'`, user_login)
            update, err2 := db.Query(query2)

            //if there is an error inserting, handle it
            if err2 != nil {
                panic(err.Error())
            }

            //Close after finishing
            defer update.Close()
        } else {

            //Prepare and perform an UPDATE query
            query2 := fmt.Sprintf(`UPDATE parkit SET qrcode=1 WHERE username='%s'`, user_login)
            update, err2 := db.Query(query2)

            //if there is an error inserting, handle it
            if err2 != nil {
                panic(err.Error())
            }

            //Close after finishing
            defer update.Close()
        }
    }

    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }

    //Return the value of rowCounts for the handler to determine login status
    return
}
