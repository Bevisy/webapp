package main

import (
	"fmt"
	"log"
	"net/http"
)

func header(resp http.ResponseWriter, req *http.Request) {
	//req.Method string
	_, err := fmt.Fprintln(resp, "Method:", req.Method)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(resp, "----------")
	if err != nil {
		log.Fatal(err)
	}

	//req.Header map[string][]string
	_, err = fmt.Fprintln(resp, "Header")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range req.Header {
		_, err = fmt.Fprintln(resp, k, ":", v)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = fmt.Fprintln(resp, "----------")
	if err != nil {
		log.Fatal(err)
	}

	//req.Context
	_, err = fmt.Fprintln(resp, req.Context())
	if err != nil {
		log.Fatal()
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", header)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
