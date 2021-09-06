package utils

import "os"

//@function: IsExists
//@description: 根据在操作文件时返回的错误信息来判断文件目录是否存在
//@param: path string
//@return: bool, error
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
