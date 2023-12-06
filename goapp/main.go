package main

import (
	"log"
	"net/http"

	"github.com/mizumoto-cn/csca5028w2/goapp/server"
)

func main() {
	http.Handle("/", server.EchoHandler{})
	http.Handle("/echo", server.EchoPrintHandler{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
