package main

import (
	"log"
	"net/http"
	"os"

	"github.com/g-e-e-z/gomicro/handlers"
)

func main() {
    l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
    hh := handlers.NewHello(l)
    gh := handlers.NewGoodbye(l)

    sm := http.NewServeMux()
    sm.Handle("/", hh)
    sm.Handle("/g", gh)

    http.ListenAndServe(":3000", sm)
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
