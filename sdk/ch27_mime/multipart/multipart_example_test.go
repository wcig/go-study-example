package multipart

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestClientWithGin(t *testing.T) {
	// 保存form表单文件目录
	const dir = "tmp"
	_ = os.Mkdir(dir, os.ModePerm)

	r := gin.Default()
	r.POST("/multipart", func(c *gin.Context) {
		// 解析form表单
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(200, gin.H{"code": -1, "msg": err.Error()})
			return
		}

		// 打印表单字符串键值对
		for k, v := range form.Value {
			fmt.Printf("string params: key:%s, val:%s\n", k, v)
		}

		// 打印表单文件信息并保存到本地
		for k, v := range form.File {
			if len(v) > 0 {
				fileHeader := v[0]
				fileName := fileHeader.Filename
				fmt.Printf("file params: fieldName:%s, fileName:%s\n", k, fileName)
				err := c.SaveUploadedFile(fileHeader, filepath.Join(dir, fileName))
				if err != nil {
					c.JSON(200, gin.H{"code": -1, "msg": err.Error()})
					return
				}
			}
		}
		c.JSON(200, gin.H{"code": 0})
	})
	r.Run(":28080") // listen and serve on 0.0.0.0:28080 (for windows "localhost:8080")

	// output:
	// string params: key:name, val:[tom]
	// string params: key:age, val:[10]
	// file params: fieldName:file2, fileName:tmp.2.txt
	// file params: fieldName:file1, fileName:tmp.1.txt
}

func TestClient(t *testing.T) {
	_ = ioutil.WriteFile("tmp.1.txt", []byte("ok"), os.ModePerm)
	_ = ioutil.WriteFile("tmp.2.txt", []byte("hello world."), os.ModePerm)

	url := "http://localhost:28080/multipart"
	strParams := map[string]string{
		"name": "tom",
		"age":  "10",
	}
	fileParams := map[string]string{
		"file1": "tmp.1.txt",
		"file2": "tmp.2.txt",
	}

	client := http.Client{
		Timeout: time.Minute,
	}
	buf := bytes.NewBuffer(nil)
	mw := multipart.NewWriter(buf)
	for k, v := range strParams {
		if err := mw.WriteField(k, v); err != nil {
			panic(err)
		}
	}
	for k, v := range fileParams {
		fw, err := mw.CreateFormFile(k, v)
		if err != nil {
			panic(err)
		}

		fr, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		// defer fr.Close()

		if _, err = io.Copy(fw, fr); err != nil {
			panic(err)
		}
	}
	if err := mw.Close(); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("response http status code:", resp.StatusCode)

	respBuf := bytes.NewBuffer(nil)
	_, err = respBuf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	fmt.Println("response body:", respBuf.String())
	// output:
	// response http status code: 200
	// response body: {"code":0}
}
