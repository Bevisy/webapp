package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	//ReadDir 列出目录的内容，返回一个 os.FileInfo 类型的切片对象
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	//Chdir 可以修改当前工作目录，类似于 cd。
	err = os.Chdir("subdir/parent/child")
	check(err)
	//列出当前目录
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	//遍历一个目录及其所有子目录。 Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

//filepath.Walk 遍历访问到每一个目录和文件后，都会调用 visit
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
