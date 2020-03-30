package main

import (
	"fmt"
	"net/http"
)

func param(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(resp, req.Form)

	//按照请求参数名获取参数值,根据源码,FormValue(key)=req.Form[key]
	name := req.FormValue("name")
	age := req.FormValue("age")
	fmt.Fprintln(resp, name, age)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/param", param)
	server.ListenAndServe()
}
