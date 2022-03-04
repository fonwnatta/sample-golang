package main

import (
	"database/sql"
	"fmt"
	"golang-training/main/internal/router"
    "golang-training/main/internal/repository"
	"golang-training/main/internal/sms"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. username:password@tcp(my-sql-db-url.com)/dbName?charset=utf8
	DB_DSN = "root:P@ssw0rd@tcp(127.0.0.1:3306)/golangtest"
)

func main() {
    fmt.Println("hello world ")

    //init connection 
    db, err := sql.Open("mysql", DB_DSN)
    if err != nil {
		fmt.Println("Cannot open DB connection", err)
	} else {
		fmt.Println("connect success")
	}
   
    
    //init repository 
    userDetailRepo := repository.NewUserDetail(db)
    

    service := sms.NewService(userDetailRepo)
    //init router
	r := router.InitRouter(service)

    //start server
    err = http.ListenAndServe(":3000", r)
    if err != nil {
        panic(err)
    }
}