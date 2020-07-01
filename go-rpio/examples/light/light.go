package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
)

var (
	pin = rpio.Pin(10)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// 设置为输出模式
	pin.Output()
	// 设置为高电平
	pin.Write(rpio.High)
}
