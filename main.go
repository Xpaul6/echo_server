package main

import (
	"log"
	"net/http"
)

const VERSION = "1.0"

func echo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.URL.Path[6:]))
	w.Write([]byte(string("\n")))
}

func GetVersion(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(string(VERSION)))
	w.Write([]byte(string("\n")))
}

func main() {
	http.HandleFunc("/echo/", echo)
	http.HandleFunc("/version", GetVersion)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
