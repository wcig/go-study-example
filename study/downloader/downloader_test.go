package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	fileUrl := "https://dlcdn.apache.org/maven/maven-3/3.8.4/binaries/apache-maven-3.8.4-bin.tar.gz"
	fileName := filepath.Base(fileUrl)

	isSupport, fileSize, err := checkSupportMultiDownload(fileUrl)
	assert.Nil(t, err)

	if isSupport {
		err = multiDownload(fileUrl, fileSize, fileName)
	} else {
		err = singleDownload(fileUrl, fileName)
	}
	assert.Nil(t, err)
}

func checkSupportMultiDownload(url string) (ok bool, fileSize int64, err error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, 0, err
	}
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return true, resp.ContentLength, nil
	}
	return false, 0, nil
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

func multiDownload(fileUrl string, fileSize int64, fileName string) error {
	dir, err := os.MkdirTemp("./", "frag-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	const fragSize int64 = 512 * 1024 // byte
	num := int(fileSize / fragSize)
	if fileSize%fragSize != 0 {
		num++
	}

	var wg sync.WaitGroup
	wg.Add(num)
	for i := 1; i <= num; i++ {
		start := int64(i-1) * fragSize
		end := int64(i)*fragSize - 1
		if end > fileSize {
			end = fileSize
		}
		n := i
		go func() {
			defer wg.Done()
			if err = downloadFrag(fileUrl, dir, n, start, end); err != nil {
				panic(err)
			}
			fmt.Printf("download frag %d over\n", n)
		}()
	}

	wg.Wait()
	return mergeFrag(fileName, dir, num)
	// return nil
}

func downloadFrag(fileUrl string, dir string, index int, start int64, end int64) error {
	fmt.Println("download frag:", fileUrl, dir, index, start, end)

	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath.Join(dir, strconv.Itoa(index)))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func mergeFrag(fileName string, dir string, num int) error {
	fmt.Println("merge frag:", fileName, dir, num)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 1; i <= num; i++ {
		fragFileName := filepath.Join(dir, strconv.Itoa(i))
		fragFile, err := os.Open(fragFileName)
		if err != nil {
			return err
		}
		defer fragFile.Close()

		if _, err = io.Copy(file, fragFile); err != nil {
			return err

		}
	}
	return nil
}

func TestDownloaderMultiDownload(t *testing.T) {
	dl := NewDownloader(5)
	fileUrl := "https://dlcdn.apache.org/maven/maven-3/3.8.4/binaries/apache-maven-3.8.4-bin.tar.gz"
	err := dl.Download(fileUrl, "")
	assert.Nil(t, err)
}
