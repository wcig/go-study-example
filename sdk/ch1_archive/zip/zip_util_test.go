package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	_ = os.Mkdir("tmp", 0755)
	_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
	_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
	_ = os.Mkdir("tmp/c", 0755)
	_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)
	defer os.RemoveAll("tmp")

	{
		// 相对路径
		err := Zip("tmp.1.zip", "tmp")
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := Zip(filepath.Join(wd, "tmp.2.zip"), filepath.Join(wd, "tmp"))
		assert.Nil(t, err)
	}
}

func TestZipV2(t *testing.T) {
	_ = os.Mkdir("tmp", 0755)
	_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
	_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
	_ = os.Mkdir("tmp/c", 0755)
	_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)
	defer os.RemoveAll("tmp")

	{
		// 相对路径
		err := ZipV2("tmp.1.zip", "tmp")
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := ZipV2(filepath.Join(wd, "tmp.2.zip"), filepath.Join(wd, "tmp"))
		assert.Nil(t, err)
	}
}

func TestUnzip(t *testing.T) {
	TestZipDir(t)

	{
		// 相对路径
		err := Unzip("", "tmp.zip")
		fmt.Println(err)
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := Unzip(filepath.Join(wd, "test"), filepath.Join(wd, "tmp.zip"))
		assert.Nil(t, err)
	}
}

/* ------------------------------------------------------------------- */

const (
	ZipFileExt = ".zip"
)

var (
	ErrEmpty  = errors.New("zip: file empty")
	ErrZipExt = errors.New("zip: file extension error")
)

// 压缩文件
func Zip(dst, src string) (err error) {
	// 校验
	if src == "" {
		return ErrEmpty
	}
	if filepath.Ext(dst) != ZipFileExt {
		return ErrZipExt
	}
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// 创建zip文件
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历并写入到zip文件
	err = filepath.Walk(fileInfo.Name(), func(path string, info fs.FileInfo, err2 error) error {
		// 遍历目录错误则直接返回
		if err2 != nil {
			return err2
		}

		// 创建zip文件头信息
		fileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 保留源文件目录结构
		fileHeader.Name = path
		if info.IsDir() {
			fileHeader.Name += "/"
		}

		// 根据zip文件头信息创建io.Writer用于写入文件
		fw, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		// 如果是目录则只写入头信息不需要写入文件数据
		// 如果不是标准文件也是一样
		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		// 拷贝数据写入
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// 压缩文件V2简化版本
func ZipV2(dst, src string) (err error) {
	// 校验
	if src == "" {
		return ErrEmpty
	}
	if filepath.Ext(dst) != ZipFileExt {
		return ErrZipExt
	}
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// 确保源文件目录为相对路径
	fileName := fileInfo.Name()

	// 创建zip文件
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 创建目录遍历函数
	walkFunc := func(path string, info fs.FileInfo, err2 error) error {
		// 错误则直接返回
		if err2 != nil {
			return err2
		}

		// 如果是目录则跳过
		if info.IsDir() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		// 创建写入zip的文件 (确保path为相对路径)
		fw, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		// 拷贝数据写入
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}

		return nil
	}

	// 执行遍历并写入zip文件
	return filepath.Walk(fileName, walkFunc)
}

// 解压文件 (如果dstDir为空则解压至当前目录)
func Unzip(dstDir, srcFile string) (err error) {
	// 校验
	if filepath.Ext(srcFile) != ZipFileExt {
		return ErrZipExt
	}

	// 读取zip文件
	readCloser, err := zip.OpenReader(srcFile)
	if err != nil {
		return err
	}

	// 目标目录不为空且不存在则创建
	if dstDir != "" {
		if !CheckFileExists(dstDir) {
			if err := os.MkdirAll(dstDir, 0755); err != nil {
				return err
			}
		}
	}

	// 遍历zip文件列表并写入
	for _, file := range readCloser.File {
		fileName := filepath.Join(dstDir, file.Name)

		// 如果是目录则创建并继续
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(fileName, file.Mode()); err != nil {
				panic(err)
			}
			continue
		}

		// 打开zip各个文件
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建写入文件
		fw, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		// 拷贝
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 关闭文件流
		fw.Close()
		fr.Close()
	}
	return nil
}

// 校验文件是否已存在
func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
