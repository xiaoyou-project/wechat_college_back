package common

import "os"

/*文件夹是否存在*/
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
	return true
}

/*创建文件夹*/
func CreatePath(filepath string) error {
	if !IsExist(filepath) {
		err := os.MkdirAll(filepath, os.ModePerm)
		return err
	}
	return nil
}
