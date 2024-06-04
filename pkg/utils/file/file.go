package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FileExists 校验文件是否存在
func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// FilesExists 校验文件是否存在
func FilesExists(names []string) bool {
	for _, name := range names {
		ok, _ := FileExists(name)
		if ok {
			return ok
		}
	}
	return false
}

// CreateFile 多级创建文件
func CreateFile(filePath string) error {
	// 获取文件所在目录的路径
	dir := filepath.Dir(filePath)

	// 递归创建目录
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("文件创建成功:", filePath)
	return nil
}
