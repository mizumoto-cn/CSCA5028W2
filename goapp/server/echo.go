package server

import (
	"log"
	"net/http"
)

// EchoHandler is a HandlerFunc that posts a form to /echo.

type EchoHandler struct{}

func (h EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("EchoHandler Called")
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	<head>
		<title>Go App</title>
	</head>
	<body>
		<h1>GoAPP:` + r.URL.Path + `</h1>
		<form action="/echo" method="POST">
			<input name="user_input">
			<input type="submit" value="Submit!">
		</form>
		<br />
		<p>Copyright@mizumoto-cn 2023</p>
		<p>This is a simple homework web application for CSCA5028W2.</p>
		<p>All the homework code is done by <a href="https://github.com/mizumoto-cn>">@mizumoto-cn</a> and follows the MIT license.</p>
	</body>
	</html>`))
}

// EchoPrintHandler is a HandlerFunc that echoes the request path together with the
// post data.
type EchoPrintHandler struct{}

func (h EchoPrintHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("EchoPrintHandler Called" + r.FormValue("user_input"))
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	<head>
		<title>Go App</title>
	</head>
	<body>
		<h1>GoAPP:` + r.URL.Path + `</h1>
		<p>` + r.FormValue("user_input") + `</p>
		<br />
		<p>Copyright@mizumoto-cn 2023</p>
		<p>This is a simple homework web application for CSCA5028W2.</p>
		<p>All the homework code is done by <a href="https://github.com/mizumoto-cn>">@mizumoto-cn</a> and follows the MIT license.</p>
	</body>
	</html>`))
}
