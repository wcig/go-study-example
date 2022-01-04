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
	concurrency          int
	fileUrl              string
	fileName             string
	fileSize             int64
	supportMultiDownload bool
	tmpDir               string
}

func NewDownloader(concurrency int) *Downloader {
	return &Downloader{concurrency: concurrency}
}

func (d *Downloader) Download(fileUrl string, fileName string) error {
	urlInfo, err := url.Parse(fileUrl)
	if err != nil {
		return err
	}

	d.fileUrl = fileUrl
	d.fileName = fileName
	if d.fileName == "" {
		d.fileName = filepath.Base(urlInfo.Path)
	}

	if err = d.isSupportMultiDownload(); err != nil {
		return err
	}
	if !d.supportMultiDownload {
		return d.singleDownload()
	} else {
		return d.multiDownload()
	}
}

func (d *Downloader) isSupportMultiDownload() error {
	resp, err := http.Head(d.fileUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		d.supportMultiDownload = true
		d.fileSize = resp.ContentLength
	}
	return nil
}

func (d *Downloader) singleDownload() error {
	resp, err := http.Get(d.fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(d.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func (d *Downloader) multiDownload() error {
	dir, err := os.MkdirTemp("./", "frag-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	d.tmpDir = dir

	var wg sync.WaitGroup
	wg.Add(d.concurrency)
	fragSize := d.fileSize / int64(d.concurrency)
	for i := 1; i <= d.concurrency; i++ {
		go func(i int, fragSize int64) {
			defer wg.Done()
			start := int64(i-1) * fragSize
			end := int64(i)*fragSize - 1
			if i == d.concurrency {
				end = d.fileSize
			}
			if err2 := d.downloadFrag(i, start, end); err2 != nil {
				err = err2
			}
			fmt.Println("download frag over:", i, start, end)
		}(i, fragSize)
	}
	if err != nil {
		return err
	}

	wg.Wait()
	return d.mergeFrag()
}

func (d *Downloader) downloadFrag(i int, start int64, end int64) error {
	fmt.Println("download frag:", i, start, end)

	req, err := http.NewRequest("GET", d.fileUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fragFileName := d.getFragFileName(i)
	file, err := os.Create(fragFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func (d *Downloader) mergeFrag() error {
	fmt.Println("merge frag")
	file, err := os.Create(d.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 1; i <= d.concurrency; i++ {
		fragFileName := d.getFragFileName(i)
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

func (d *Downloader) getFragFileName(i int) string {
	return filepath.Join(d.tmpDir, fmt.Sprintf("frag-%d", i))
}
