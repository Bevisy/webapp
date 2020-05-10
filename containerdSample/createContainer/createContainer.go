package main

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
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
	ctx := namespaces.WithNamespace(context.Background(), "example")
	redisImage := "docker.io/library/redis:alpine"
	image, err := client.Pull(ctx, redisImage, containerd.WithPullUnpack)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully pulled %s image\n", image.Name())
}
