package main

import (
	"net/http"
	"strconv"
)

// simple post handler
// posts given string
type PostHandler struct {
	PostString string
}

func (h PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.PostString))
}

func main() {
	// create a new mux
	mux := http.NewServeMux()
	count := 0
	// create a json string
	jsonStr := `{"name":"mizumoto-cn","count":` + strconv.Itoa(count) + `}`
	// register handlers
	mux.Handle("/data", PostHandler{PostString: jsonStr})

	// start server
	http.ListenAndServe(":8081", mux)
}
