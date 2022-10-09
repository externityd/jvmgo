package classpath

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs}
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//1. 尝试读取zip文件
	r, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r)

	//2. 遍历zip文件，尝试找到类
	for _, f := range r.File {
		// 如果文件名等于className则返回Entry信息,否则返回异常
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer func(rc io.ReadCloser) {
				err := rc.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(rc)
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, entry, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}
