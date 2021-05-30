package tar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// archive/tar
// tar gzip压缩解压相关

func TestTarGzipFiles(t *testing.T) {
	{
		_ = os.WriteFile("tmp.a.txt", []byte("aaaaa"), 0755)
		_ = os.WriteFile("tmp.b.txt", []byte("bbbbb"), 0755)
		_ = os.WriteFile("tmp.c.txt", []byte("ccccc"), 0755)

		err := TarGzipFiles("tmp.1.tar.gz", []string{"tmp.a.txt", "tmp.b.txt", "tmp.c.txt"})
		assert.Nil(t, err)
	}
	{
		dir := "/Users/yangbo/Desktop/"
		_ = os.WriteFile(filepath.Join(dir, "tmp.a.txt"), []byte("aaaaa"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.b.txt"), []byte("bbbbb"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.c.txt"), []byte("ccccc"), 0755)

		err := TarGzipFiles(filepath.Join(dir, "tmp.1.tar.gz"),
			[]string{
				filepath.Join(dir, "tmp.a.txt"),
				filepath.Join(dir, "tmp.b.txt"),
				filepath.Join(dir, "tmp.c.txt"),
			})
		assert.Nil(t, err)
	}
}

func TestTarGzipDir(t *testing.T) {
	{
		_ = os.Mkdir("tmp", 0755)
		_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
		_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
		_ = os.Mkdir("tmp/c", 0755)
		_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)

		err := TarDir("tmp.2.tar.gz", "tmp")
		assert.Nil(t, err)
	}
	{
		dir := "/Users/yangbo/Desktop/"
		_ = os.Mkdir(filepath.Join(dir, "tmp"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp/a"), []byte("aaaaa"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp/b"), []byte("bbbbb"), 0755)
		_ = os.Mkdir(filepath.Join(dir, "tmp/c"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp/c/d"), []byte("ddddd"), 0755)

		err := TarDir(filepath.Join(dir, "tmp.2.tar.gz"), filepath.Join(dir, "tmp"))
		assert.Nil(t, err)
	}
}

func TestUnTarGzip(t *testing.T) {
	{
		err := UnTar("", "tmp.2.tar.gz")
		assert.Nil(t, err)
	}
	{
		dir := "/Users/yangbo/Desktop/"
		err := UnTar(dir, filepath.Join(dir, "tmp.2.tar.gz"))
		assert.Nil(t, err)
	}
}

/* ------------------------------------------------------------------- */

var (
	ErrEmpty2 = errors.New("tar gzip: file empty")
)

// 压缩多个文件
func TarGzipFiles(dst string, files []string) (err error) {
	if len(files) == 0 {
		return ErrEmpty2
	}

	tarFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, fileName := range files {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			panic(err)
		}

		header, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			panic(err)
		}
		header.Name = fileInfo.Name()

		if err = tarWriter.WriteHeader(header); err != nil {
			panic(err)
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()

		if _, err = io.Copy(tarWriter, fr); err != nil {
			return err
		}
	}

	return nil
}

// 压缩单个目录
func TarGzipDir(dst string, src string) (err error) {
	if src == "" || dst == "" {
		return ErrEmpty2
	}

	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	baseDir := fileInfo.Name()

	tarFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	walkFunc := func(path string, info fs.FileInfo, err2 error) error {
		if err2 != nil {
			return err2
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, src))

		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		_, err = io.Copy(tarWriter, fr)
		return err
	}

	err = filepath.Walk(src, walkFunc)
	return err
}

// 解压文件 (dst为空则解压到当前目录)
func UnTarGzip(dst string, src string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	gzipReader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		fileName := filepath.Join(dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(fileName, header.FileInfo().Mode())
			if err != nil {
				return err
			}
		case tar.TypeReg:
			fw, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, header.FileInfo().Mode())
			if err != nil {
				return err
			}

			if _, err = io.Copy(fw, tarReader); err != nil {
				fw.Close()
				return err
			}
			fw.Close()
		}
	}
}
