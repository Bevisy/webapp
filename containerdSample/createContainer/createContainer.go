package main

import (
	"github.com/containerd/containerd"
	"log"
)

func redisExample() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}

	defer client.Close()
	return nil
}
func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}
