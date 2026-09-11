package main

import (
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.URL.Path))
	w.Write([]byte(string("\n")))
}

func main() {
	http.HandleFunc("/", echo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
