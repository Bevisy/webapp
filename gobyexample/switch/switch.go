package main

import (
	"fmt"
	"time"
)

func main() {

	//基本的 switch
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}

	//同一个case中使用逗号分隔多个表达式
	//default 分支 可选
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("have a good rest :)")
	default:
		fmt.Println("job day :(")
	}

	//不带表达式的switch是实现if/else逻辑的另一种方式。
	//这里还展示了case表达式可以不使用常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	//类型开关(type switch)比较类型而非值。可以用来发现一个接口值的类型。
	//这个例子中，变量t在每个分支中会有相应的类型
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("i m bool")
		case string:
			fmt.Println("i m string")
		default:
			fmt.Println("none of your business")
		}
	}
	whatAmI(true)
	whatAmI("sss")
}
