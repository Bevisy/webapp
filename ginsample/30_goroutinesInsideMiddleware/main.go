package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//When starting new Goroutines inside a middleware or handler,
//you SHOULD NOT use the original context inside it, you have to use a read-only copy.
func main() {
	r := gin.Default()
	r.GET("/async", func(c *gin.Context) {
		//Copy returns a copy of the current context that can be safely used outside the request's scope.
		//This has to be used when the context has to be passed to a goroutine.
		cCp := c.Copy() //获取一份只读拷贝
		//请求立即返回，日志在请求返回后输出
		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("async done! The request is %s.\n", cCp.Request.URL.Path)
		}()
	})

	r.GET("/sync", func(c *gin.Context) {
		//执行如下语句并输出日志后，请求返回
		time.Sleep(5 * time.Second)
		log.Printf("sync done! The request is %s.\n", c.Request.URL.Path)
	})
	r.Run("127.0.0.1:8080")
}
