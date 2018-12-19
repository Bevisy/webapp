package main

import (
	"log"
	"net/http"
	"os"
	"syscall"
)

func echobot(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, req.URL.Path+"\n")
	_, err := w.Write([]byte(req.URL.Path))
	if err != nil {
		log.Fatalln("REPONSE FAILED!\n")
	}
	log.Fatalln("REQUEST URL: ", req.URL.Path)
}

func main() {
	log.SetFlags(syscall.Stdin)
	http.HandleFunc("/", echobot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("SERVER FAILED!\n")
		os.Exit(-1)
	}
}
