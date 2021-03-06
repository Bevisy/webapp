package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

//参数2的值，拷贝到参数1
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

var host = flag.String("host", "127.0.0.1", "the remote server host")
var port = flag.Int("port", 8000, "the remote server port")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
