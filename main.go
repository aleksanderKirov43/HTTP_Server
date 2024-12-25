package main

import "net/http"

func main() {
	err := http.ListenAndServe(`:8088`, nil)
	if err != nil {
		panic(err)
	}
}
