package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 需要处理的JSON文档以string的形式存在
// 首先需要将string转换成byte切片([]byte), 再使用json包的Unmarshal函数进行反序列化处理

// Contact 结构代表我们的JSON字符串
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON 包含用于反序列化的演示字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415",
		"cell": "415"
	}
}`

func main() {
	// 将JSON字符串反序列化到变量
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v", c)
}
