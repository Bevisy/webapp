package main

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"log"
)

//func listImages(client *containerd.Client) error{
//	client.Containers()
//}

func main() {
	path := "/run/containerd/containerd.sock"
	//新建containerd client
	client, err := containerd.New(path)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	//指定 containerd 将要使用的命名空间
	ctx := namespaces.WithNamespace(context.Background(), "example")
	cImage := "docker.io/library/redis:alpine"
	//拉取指定容器镜像，将其解压缩到快照程序中，以用作根文件系统。
	//（unpack it into a snapshotter for use as a root filesystem.）
	image, err := client.Pull(ctx, cImage, containerd.WithPullUnpack)
	if err != nil {
		panic(err)
	}
	//image.Name() 示例：docker.io/library/busybox:latest
	log.Printf("Successfully pulled %s image\n", image.Name())

	//创建容器
	container, err := client.NewContainer(ctx, "redis-server",
		containerd.WithNewSnapshot("redis-server-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
	)
	if err != nil {
		panic(err)
	}
	log.Printf("container %s create succeess.", container.ID())
	//创建成功后即删除
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
}
