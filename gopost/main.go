package main

import (
	"log"
	"net/http"
	"strconv"
)

var count int = 0

// simple post handler
// posts given string
type PostHandler struct {
	PostString string
}

func (h PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.PostString))
	log.Println("POST::", r.URL.Path, "::", h.PostString)
	count++
}

func main() {
	// create a new mux
	mux := http.NewServeMux()
	// create a json string
	jsonStr := `{"name":"mizumoto-cn","count":` + strconv.Itoa(count) + `}`
	// register handlers
	mux.Handle("/data", PostHandler{PostString: jsonStr})

	// start server
	http.ListenAndServe(":8081", mux)
}
