package main

import (
	"github.com/containerd/containerd"
	"log"
)

//func listImages(client *containerd.Client) error{
//	client.Containers()
//}

func main() {
	path := "/run/containerd/containerd.sock"
	client, err := containerd.New(path)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Run success.")
}
