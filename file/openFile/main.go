package main

import "os"

// 大文件读写 https://www.tizi365.com/archives/337.html

// func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

/**
 * name - 文件路径
 * flag - 文件打开模式
 * perm - 文件权限模式
 */

/**
 * 常用的文件打开模式,可以组合
 * O_CREATE - 如果文件不存在,则创建一个
 * O_APPEND - 写入文件的内容,自动追加到文件的尾部
 * O_RDONLY - 打开一个只读的文件
 * O_WRONLY - 打开一个只写的文件
 * O_RDWR - 打开一个可以读写的文件
 */

func main() {
	// 打开文件并设置文件内容以追加的形式添加到文件尾部, 如果文件不存在则创建一个
	f, err := os.OpenFile("./demo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close() // 读写结束后关闭文件

	f.WriteString("内容1\n")
	f.WriteString("内容2\n")
	f.WriteString("内容3\n")
}
