package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
			<title>Olá mundo!</title>
			</head>
			<body>
				Ola, mundo!
			</body>
		</html>`)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)

}
