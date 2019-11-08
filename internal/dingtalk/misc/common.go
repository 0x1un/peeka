package misc

import "os"

// 这个文件包含一些此API常用的必要组件

func IsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
