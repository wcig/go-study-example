package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	fileUrl := "https://dlcdn.apache.org/maven/maven-3/3.8.4/binaries/apache-maven-3.8.4-bin.tar.gz"
	fileName := "apache-maven-3.8.4-bin.tar.gz"

	err := Download(fileUrl, fileName)
	assert.Nil(t, err)

	err = NewDownloader(5).Download(fileUrl, fileName+".2")
	assert.Nil(t, err)
}
