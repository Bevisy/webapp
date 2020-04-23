package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

// curl -X POST http://127.0.0.1:8080/upload \
//-F "files=@test.txt" \
//-F "files=@test2.txt" \
//-F "files=@test3.txt \
//-H "Content-Type: multipart/form-data"

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/", "./public")

	r.POST("/upload", func(c *gin.Context) {
		//multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s",
				err.Error()))
		}

		files := form.File["files"] //获取文件名称

		for _, file := range files {
			filename := "public/" + filepath.Base(file.Filename)
			log.Printf("upload file: %s", filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest,
					fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusBadRequest,
			fmt.Sprintf("upload successfully %d files", len(files)))
	})

	r.Run(":8080")
}
