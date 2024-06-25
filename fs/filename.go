package fs

import (
	"os"
	"path/filepath"
	"strings"
)

/*
返回文件名 不包含后缀 对于需要文件名包括后缀 应该使用filepath.Base(src)方法
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(src string) string {
	ext := filepath.Ext(src)
	src = filepath.Base(src)
	length := len(src)
	return src[:length-len(ext)]
}

/*
返回文件的文件名 携带前缀的路径 并且一定为绝对路径

例如：/home/xxx/xxx.go 返回 /home/xxx/xxx

test.go 返回 /home/xxx/test

对于path.Dir()方法 区别就是可以对于相对路径也返回它的绝对路径的文件夹路径
需要文件名包括后缀 应该使用filepath.Base(src)方法
*/
func FullFileName(src string) string {
	if filepath.IsAbs(src) {
		return src[:len(src)-len(filepath.Ext(src))]
	} else {
		abs_src, _ := filepath.Abs(src)
		return abs_src[:len(abs_src)-len(filepath.Ext(abs_src))]
	}
}

/*
判断文件或者文件夹是否存在
*/
func IsExit(src string) bool {
	_, err := os.Stat(src)
	return err != nil
}

/*
将文件类型换为特定类型的文件类型 返回相对路径 不需要带 . 后缀
例如 ChangeFileType("/home/xxx/xxx.go","txt") 返回 /home/xxx/xxx.txt
ChangeFileType("xxx","mp4") 返回 xxx.mp4
*/
func ChangeFileType(path string, type_ string) string {
	ext := filepath.Ext(path)
	withExt := strings.TrimSuffix(path, ext)
	return withExt + "." + type_
}

/*
检查文件类型是否为特定类型
例如 main.go go true
main.go txt false
*/
func IsFileType(file string, type_ string) bool {
	return filepath.Ext(file) == "."+type_
}
