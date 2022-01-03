package downloader

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	fileUrl := "https://dlcdn.apache.org/maven/maven-3/3.8.4/binaries/apache-maven-3.8.4-bin.tar.gz"
	fileName := filepath.Base(fileUrl)

	isSupport, err := checkSupportMultiDownload(fileUrl)
	assert.Nil(t, err)

	if isSupport {
		err = multiDownload(fileUrl, fileName)
	} else {
		err = singleDownload(fileUrl, fileName)
	}
	assert.Nil(t, err)
}

func checkSupportMultiDownload(url string) (ok bool, err error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return true, nil
	}
	return false, nil
}

func singleDownload(fileUrl string, fileName string) error {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func multiDownload(fileUrl string, fileName string) error {
	// 1.get file content length by head request
	// 2.send multi download request
	// 3.merge multi file
}
