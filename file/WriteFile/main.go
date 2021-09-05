package main

import "io/ioutil"

// 向指定的文件写入数据, 如果文件不存在将按给出的权限创建文件, 存在则用写入的数据覆盖已存在内容

func main() {
	content := "文件内容"

	// 权限部分参考linux文件权限相关内容
	// https://www.runoob.com/linux/linux-file-attr-permission.html
	err := ioutil.WriteFile("./demo.txt", []byte(content), 0666)

	if err != nil {
		panic(err)
	}
}
