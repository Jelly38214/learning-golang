package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	result struct {
		Group   string   `json:"group"`
		Persons []person `json:"persons"`
	}
)

func main() {
	// uri := "https://randomuser.me/api/?results=1"
	var data result
	bytes, _ := ioutil.ReadFile("data.json")
	fmt.Println("*** data.json content: ***")

	// 打印时需要转为字符串
	fmt.Println(string(bytes))

	// 将字节切片映射到指定结构上
	json.Unmarshal(bytes, &data)

	fmt.Println("*** unmarshal results: ***")

	// 打印对象结构
	fmt.Println(data)

	// 更改数据
	data.Group = "engineer"

	// 将更改后的结构对象序列化成JSON格式
	newBytes, _ := json.Marshal(&data)

	fmt.Println("*** update content: ***")

	// 打印JSON结果
	fmt.Println(string(newBytes))

}
