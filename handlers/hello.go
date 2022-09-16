package handlers

import (
    "fmt"
    "io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
    l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
    return &Hello{l}
}

// Method that satisies the http interface
func (h*Hello) ServeHTTP( rw http.ResponseWriter, r *http.Request) {
    h.l.Println("Hello world")

    d, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(rw, "Oops", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(rw, "Hello %s\n", d)
}


/*
Dont want to create concrete objects inside of a handler

*/
