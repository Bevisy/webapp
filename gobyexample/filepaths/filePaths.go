package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	//应使用 Join 来构建可移植(跨操作系统)的路径。 它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	//应该总是使用 Join 代替手动拼接 / 和 \。 除了可移植性，Join 还会删除多余的分隔符和目录，使得路径更加规范。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	//判断路径是否是绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	//文件后缀
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	//文件名去除后缀
	fmt.Println(strings.TrimSuffix(filename, ext))

	//Rel 寻找 basepath 与 targpath 之间的相对路径。 如果相对路径不存在，则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
