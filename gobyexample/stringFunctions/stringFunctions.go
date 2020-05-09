package main

import (
	"fmt"
	s "strings" // 添加别名
)

//fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

func main() {

	//这是一些 strings 中有用的函数例子。 由于它们都是包的函数，而不是字符串对象自身的方法，
	//这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。 你可以在 strings 包文档中找到更多的函数。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
