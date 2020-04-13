package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//s, sep := "", ""
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
	for inx, arg := range os.Args[0:] {
		//s += sep + arg
		//sep = " "
		fmt.Println(inx, arg)
	}
	//fmt.Println(s)
}
