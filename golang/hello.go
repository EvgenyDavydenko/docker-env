package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!!!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    log.Println("Server starting on port 5050...")
    log.Fatal(http.ListenAndServe("golang:5050", nil))
}


