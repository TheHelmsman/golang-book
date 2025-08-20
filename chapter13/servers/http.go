package main

import (
	"fmt"
	"net/http"
	"io"
)

/*
	Open browser
	Write following URI: `http://127.0.0.1:9000/hello`

	`HandleFunc` handles a URL route (`/hello`) by calling the given function. 
	We can also handle static files by using `FileServer`:

		http.Handle(
			"/assets/",
			http.StripPrefix(
				"/assets/",
				http.FileServer(http.Dir("assets")),
			),
		)
	
*/

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
			<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				Hello World!
			</body>
		</html>`,
	) 
}

func main() {
	fmt.Println("Running HTTP server on 127.0.0.1 port 9000")
	fmt.Println("Start serving resources:")
	fmt.Println("- /hello")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("127.0.0.1:9000", nil)
}