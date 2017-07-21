package main

import (
    "net/http"
	"log"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)
    r.HandleFunc("/home", Home)

    log.Fatal(http.ListenAndServe(":8081", nil))
}