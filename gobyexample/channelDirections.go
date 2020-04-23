package main

import "fmt"

func ping(pings chan<- string, msg string) { //写入信息到信道
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { //读取信息并写入新的信道
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
