package mapknowledge

import "fmt"

// 應說明 exists 布林值的用途，以及為什麼不能只用 value != nil 來判斷。
func MapExist() {
	myMap := map[string]int{"apple": 1, "banana": 2}
	value, exists := myMap["apple"]

	fmt.Sprint(value)
	fmt.Sprint(exists)
	// 不能用 value != nil 來判斷的原因是因為如果不存在，在 value 會給該類別的零值
	// int => 0 , string => "", map, slice, chan => nil
}
