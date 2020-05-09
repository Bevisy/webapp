package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//通过 Compile 得到一个优化过的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	//创建正则表达式常量时，可以使用 Compile 的变体 MustCompile。
	//普通的 Compile 不适用于常量，因为它有 2 个返回值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
