package oss

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	endpoint   = "https://oss-cn-shenzhen.aliyuncs.com"
	bucketName = "example2024"
)

func TestSimpleUploadFile(t *testing.T) {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		log.Fatal(err)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(endpoint, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		log.Fatal(err)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Fatal(err)
	}

	// 计算md5
	fileName := "tmp/example.txt"
	md5Val, err := FileMD5(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(fileName, fileName, oss.ContentMD5(md5Val))
	if err != nil {
		log.Fatal(err)
	}
}

// 计算文件MD5
func FileMD5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := md5.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
