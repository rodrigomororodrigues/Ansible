package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Essa é a Homepage Desafio  \n</h1>")
    })

    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Desafio Serasa - RodrigoMoro!\n</h1>")
    })

    r.HandleFunc("/ola/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["name"]

        fmt.Fprintf(w, "<h1>Ola, %s!\n</h1>", title)
    })

    http.ListenAndServe(":8080", r)
}
