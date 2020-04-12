package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func echo(c net.Conn, s string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		//go(routine) 后跟的函数的参数会在go语句自身执行时被求值
		go echo(c, input.Text(), 2*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

var host = flag.String("host", "127.0.0.1", "the server listening host")
var port = flag.Int("port", 8000, "the server listening port")

func main() {
	//把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	//Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接
	listener, err := net.Listen("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
