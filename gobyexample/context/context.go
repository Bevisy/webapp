package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println("server: hello handler started.")
	defer log.Println("server: hello handler ended.")

	stop := time.NewTimer(10 * time.Second)
	select {
	case <-stop.C:
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		log.Println("server: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
