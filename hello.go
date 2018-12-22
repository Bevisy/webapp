package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		log.Fatalln("ParseForm() failed!")
	}
	//fmt.Println(r.Form)
	//fmt.Println("path: ", r.URL.Path)
	//fmt.Println("scheme: ", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//    fmt.Println("key:", k, "\tval:", strings.Join(v, ""))
	//}
	_, err = fmt.Fprintf(w, "Hello!\n")
	if err != nil {
		log.Fatalln("Write message to client failed!")
	}
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
