package main

import "net/http"

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", handler)
	http.ListenAndServe(":8080", myMux)
}

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, world"))
}
