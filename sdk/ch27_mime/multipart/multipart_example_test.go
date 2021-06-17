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
	"gopkg.in/resty.v1"
)

func TestServerWithGin(t *testing.T) {
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

func TestClientWithResty(t *testing.T) {
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

	client := resty.New()
	resp, err := client.R().
		SetFormData(strParams).
		SetFiles(fileParams).
		Post(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())
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

	// 创建http client
	client := http.Client{
		Timeout: time.Minute,
	}
	// 创建buf的io.Reader保存所有表单数据
	buf := bytes.NewBuffer(nil)
	// 创建multipart.Writer
	mw := multipart.NewWriter(buf)

	// 写入字符串
	for k, v := range strParams {
		if err := mw.WriteField(k, v); err != nil {
			panic(err)
		}
	}

	// 写入文件
	for k, v := range fileParams {
		fw, err := mw.CreateFormFile(k, v)
		if err != nil {
			panic(err)
		}

		fr, err := os.Open(v)
		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(fw, fr); err != nil {
			panic(err)
		}
		fr.Close()
	}

	// 注意关闭multipart.Writer,不关闭将影响边界
	if err := mw.Close(); err != nil {
		panic(err)
	}

	// 基于url和表单数据创建http.Request
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", mw.FormDataContentType())

	// 发起http请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("response http status code:", resp.StatusCode)

	// 保存http请求响应
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

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: ":28080",
	}
	http.HandleFunc("/multipart", multipartHandlerFunc)
	server.ListenAndServe()
}

func multipartHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// 保存form表单文件目录
	const dir = "tmp"
	_ = os.Mkdir(dir, os.ModePerm)

	if err := r.ParseMultipartForm(0); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	form := r.MultipartForm

	// 打印表单字符串键值对
	// 方式一
	for k, v := range form.Value {
		fmt.Printf("string params: key:%s, val:%s\n", k, v)
	}
	// 方式二
	for k, v := range r.PostForm {
		fmt.Printf("string params: key:%s, val:%s\n", k, v)
	}

	// 打印表单文件信息并保存到本地
	for k, v := range form.File {
		if len(v) > 0 {
			fileHeader := v[0]
			fileName := fileHeader.Filename
			fmt.Printf("file params: fieldName:%s, fileName:%s\n", k, fileName)
			err := SaveUploadedFile(fileHeader, filepath.Join(dir, fileName))
			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}
		}
	}
	fmt.Fprintf(w, "ok")
	// output:
	// string params: key:name, val:[tom]
	// string params: key:age, val:[10]
	// string params: key:name, val:[tom]
	// string params: key:age, val:[10]
	// file params: fieldName:file1, fileName:tmp.1.txt
	// file params: fieldName:file2, fileName:tmp.2.txt
}

// 保存文件到本地
func SaveUploadedFile(fileHeader *multipart.FileHeader, dst string) error {
	fr, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer fr.Close()

	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return err
}
