package handlers

import (
	"log"
	"net/http"
)

// Structure

type Goodbye struct {
    l *log.Logger
}

// Constructor of the Structure
func NewGoodbye(l *log.Logger) *Goodbye {
    return &Goodbye{l}
}

// Create a Handler with a ServeHTTP method to adhere to interface
// Add the method TO the struct

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    rw.Write([]byte("Goodbye \n"))
}
