package main

import "fmt"

func main() {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#229b22",
	}

	fmt.Printf("%X", &colors)

	// 判断是否存在某个key
	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key Blue not exists")
	}

	// 迭代map
	for key, value := range colors {
		fmt.Printf("Key: %s Value : %s\n", key, value)
	}

	// 删除某个key
	// 在函数间传递映射,并不会创建一个副本
	removeColor(colors, "Coral")

	for key, value := range colors {
		fmt.Printf("Key: %s Value : %s\n", key, value)
	}

}

func removeColor(colors map[string]string, key string) {
	fmt.Printf("%X", &colors)
	delete(colors, key)
}
