package main

import (
	"encoding/json"
	"fmt"
)

// 有时,无法为JSON的格式声明一个结构类型,而是需要更加灵活的方式来处理JSON文档
// 这种情况, 可以将JSON文档解码到一个map中

// JSON 内容
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415",
		"cell": "415"
	}
}`

func main() {
	// 将JSON字符串反序列化到map变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Tile:", c["title"])

	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"]) // 需要对interface{}做类型转换
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])

}
