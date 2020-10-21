package main

import (
	"log"
	"net/http"
)

func main() {

	h := http.FileServer(http.Dir("."))
	p := ":8000"

	log.Println("Serving at" + p)

	http.ListenAndServe(p, h)
}
