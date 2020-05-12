package main

import (
	"context"
	"fmt"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"log"
	"syscall"
	"time"
)

func redisExample() error {
	//create new client connected to the default socket path for containerd
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	//create a new context with an "example" namespace
	ctx := namespaces.WithNamespace(context.Background(), "example")
	//pull the redis image from DockerHub
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine",
		containerd.WithPullUnpack)
	if err != nil {
		return err
	}

	log.Printf("successfully pulled %s \n.", image.Name())
	//create a container
	container, err := client.NewContainer(ctx, "redis-server",
		containerd.WithNewSnapshot("redis-server-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
		containerd.WithRuntime("io.containerd.runtime.v1.linux", nil))
	if err != nil {
		return err
	}
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
	log.Printf("successfully created container with ID %s and snapshot with"+
		" ID redis-server-snapshot", container.ID())

	//create a task from container
	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}
	defer task.Delete(ctx)

	//make sure we wait before calling start
	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		return err
	}

	//call start on the task to execute the redis server
	if err := task.Start(ctx); err != nil {
		return err
	}

	//sleep for a lil bit to see the logs
	time.Sleep(3 * time.Second)

	//kill the process and get the exit status
	if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
		return err
	}

	//wait for the process to fully exit and print out the exit status

	status := <-exitStatusC
	code, _, err := status.Result()
	if err != nil {
		return err
	}

	fmt.Printf("redis-server exited with status: %d\n", code)

	return nil
}

func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}
