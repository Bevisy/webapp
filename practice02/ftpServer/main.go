package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	os.Mkdir("file", 0777)

	http.Handle("/bevisy/", http.StripPrefix("/bevisy/", http.FileServer(http.Dir("file"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server: ", err)
	}
}
