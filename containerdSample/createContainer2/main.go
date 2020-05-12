package main

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"log"
)

func redisExample() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "example")
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine",
		containerd.WithPullUnpack)
	if err != nil {
		return err
	}

	log.Printf("successfully pulled %s \n.", image.Name())

	container, err := client.NewContainer(ctx, "redis-server",
		containerd.WithNewSnapshot("redis-server-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)))
	if err != nil {
		return err
	}
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
	log.Printf("successfully created container with ID %s and snapshot with"+
		" ID redis-server-snapshot", container.ID())
	return nil
}

func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}
