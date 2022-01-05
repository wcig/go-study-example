package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

// 支持并发下载的下载器（参考：https://polarisxu.studygolang.com/posts/go/action/build-a-concurrent-file-downloader/）

type Downloader struct {
	concurrency int
}

func NewDownloader(concurrency int) *Downloader {
	return &Downloader{concurrency: concurrency}
}

var DefaultDownloader = NewDownloader(3)

func Download(fileUrl string, fileName string) error {
	return DefaultDownloader.Download(fileUrl, fileName)
}

func (d *Downloader) Download(fileUrl string, fileName string) error {
	saveFileName, err := d.getFileName(fileUrl, fileName)
	if err != nil {
		return err
	}

	isSupport, fileSize, err := d.isSupportMultiDownload(fileUrl)
	if err != nil {
		return err
	}

	if isSupport {
		return d.multiDownload(fileUrl, saveFileName, fileSize)
	}
	return d.singleDownload(fileUrl, saveFileName)
}

func (d *Downloader) getFileName(fileUrl string, fileName string) (string, error) {
	if fileName != "" {
		return fileName, nil
	}

	urlInfo, err := url.Parse(fileUrl)
	if err != nil {
		return "", err
	}

	return filepath.Base(urlInfo.Path), nil
}

func (d *Downloader) isSupportMultiDownload(fileUrl string) (isSupport bool, fileSize int64, err error) {
	resp, err := http.Head(fileUrl)
	if err != nil {
		return false, 0, err
	}
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		isSupport = true
		fileSize = resp.ContentLength
	}
	return isSupport, fileSize, nil
}

func (d *Downloader) singleDownload(fileUrl string, fileName string) error {
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

func (d *Downloader) multiDownload(fileUrl string, fileName string, fileSize int64) error {
	fragDir, err := os.MkdirTemp("./", fmt.Sprintf("%s-", fileName))
	if err != nil {
		return err
	}
	defer os.RemoveAll(fragDir)

	fragNum := d.concurrency
	fragSize := fileSize / int64(fragNum)

	var wg sync.WaitGroup
	wg.Add(fragNum)

	for i := 1; i <= fragNum; i++ {
		go func(i int, fragSize int64) {
			defer wg.Done()

			start := int64(i-1) * fragSize
			end := int64(i)*fragSize - 1
			if i == fragNum {
				end = fileSize
			}
			if err2 := d.downloadFrag(fileUrl, fragDir, i, start, end); err2 != nil {
				err = err2
			}
			fmt.Println("download frag over:", i, start, end)
		}(i, fragSize)
	}
	if err != nil {
		return err
	}

	wg.Wait()
	return d.mergeFrag(fragDir, fileName)
}

func (d *Downloader) downloadFrag(fileUrl string, fragDir string, i int, start int64, end int64) error {
	fmt.Println("download frag:", i, start, end)

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

	fragFileName := d.getFragFileName(fragDir, i)
	file, err := os.Create(fragFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func (d *Downloader) mergeFrag(fragDir string, fileName string) error {
	fmt.Println("merge frag start:")

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 1; i <= d.concurrency; i++ {
		var (
			fragFileName string
			fragFile     *os.File
		)

		fragFileName = d.getFragFileName(fragDir, i)
		fragFile, err = os.Open(fragFileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(file, fragFile)
		fragFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Downloader) getFragFileName(fragDir string, i int) string {
	return filepath.Join(fragDir, fmt.Sprintf("frag-%d", i))
}
