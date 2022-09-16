package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    log.Println("Hello world")
    d, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Oops", http.StatusBadRequest)
        return
    }

    log.Printf("Data %s\n", d)

    fmt.Fprintf(w, "Hello %s\n", d)
}

func bye(w http.ResponseWriter, r *http.Request) {
    log.Println("Goodbye world")
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/g", bye)
    http.ListenAndServe(":3000", nil)
}

/*
http.HandleFunc: Conveinience method on go http
    -> Registers a function to a path on the `Default Serve Mux`
    Default Serve Mux: An httphandler
        Everything relating to http server in go is an http handler

http.ListenAndServe: Conveinience method
    -> Construct an http server
        1st Param: register a default handler
        2md Param: Handler, when nil, uses default serve Mux






*/
