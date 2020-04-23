package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// upload
//curl -X POST http://127.0.0.1:8080/upload -F "file=@test.txt" -H "Content-Type: multipart/form-data"

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/", "./public")

	r.POST("/upload", func(c *gin.Context) {
		//FormFile
		file, err := c.FormFile("file") //请求文件名称
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s",
				err.Error()))
		}

		//拼接上传路径; filepath.Base 获取文件名
		filename := "public/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest,
				fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK,
			fmt.Sprintf("File %s uploaded successfully", file.Filename))
	})

	r.Run(":8080")
}
