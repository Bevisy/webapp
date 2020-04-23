package main

import (
	"log"
	"time"
)

func tickerDemo() {
	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func main() {
	tickerDemo()
	log.Println("end.")
}
