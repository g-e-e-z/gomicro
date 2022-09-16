package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/g-e-e-z/gomicro/handlers"
)

func main() {
    l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
    ph := handlers.NewProducts(l)

    sm := http.NewServeMux()
    sm.Handle("/", ph)

    s := &http.Server{
        Addr: ":3000",
        Handler: sm,
        IdleTimeout: 120*time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
    }

    go func() {
        err := s.ListenAndServe()
        if err != nil {
            l.Fatal(err)
        }
    }()

    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <- sigChan
    l.Println("Recieved terminate, graceful shutdown", sig)
    tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)

    s.Shutdown(tc)
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
