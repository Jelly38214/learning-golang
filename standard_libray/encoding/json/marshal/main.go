package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 序列化是指将数据转换为[]byte的过程
func main() {
	// 创建一个保存键值对的映射
	c := make(map[string]interface{})

	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "134",
		"cell": "134",
	}

	// 将这个映射序列化到JSON字符串
	data, err := json.MarshalIndent(c, "", "  ")
	data2, err2 := json.Marshal(c) // 没有缩进和格式化

	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	if err2 != nil {
		log.Println("ERROR: ", err2)
		return
	}

	fmt.Println(string(data))
	fmt.Println(string(data2))
}
