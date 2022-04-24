package util

import (
	"io/fs"
	"io/ioutil"
)

// 一次性读取文件
func ReadFileOnce(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// 写文件
func WriteFile(filePath string, fileBty []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filePath, fileBty, perm)
}
