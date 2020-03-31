package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func header(resp http.ResponseWriter, req *http.Request) {
	//req.Method string
	_, err := fmt.Fprintln(resp, "Method:", req.Method+"\n==========")
	if err != nil {
		log.Fatal(err)
	}

	//req.Header map[string][]string
	_, err = fmt.Fprintln(resp, "Header:")
	for k, v := range req.Header {
		_, err = fmt.Fprintln(resp, k, ":", v)
		if err != nil {
			log.Fatal(err)
		}
	}

	// req.Body
	body, _ := ioutil.ReadAll(req.Body)
	_, err = fmt.Fprintln(resp, "==========\n", "Body: \n", string(body))
	if err != nil {
		log.Fatal(err)
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
