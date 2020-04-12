package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		if _, err := io.WriteString(c, time.Now().Format("15:04:05\n")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

var host = flag.String("host", "127.0.0.1", "the listening host")
var port = flag.Int("port", 8000, "the listening port")

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
