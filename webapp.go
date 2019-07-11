package main

import (
    "io"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/hello", hello)
    err := http.ListenAndServe(":80", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
