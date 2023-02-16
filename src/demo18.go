package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	v1 := m["apple"]
	v2, ok := m["pear"]
	// v2不存在 由于时int 所以值为该类型的零值
	// 键不存在，ok为false
	fmt.Println(v1, v2, ok)
	//m := make(map[string]int, 10)
	for k, v := range m {
		fmt.Println(k, v)
		println()
	}
	main12()
}

func main12() {
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["google"] = "谷歌"
	siteMap["baidu"] = "百度"
	for k, v := range siteMap {
		fmt.Println(k, v)
		fmt.Println("")
	}
	delete(siteMap, "baidu")
	for k, v := range siteMap {
		fmt.Println("baidu被删除")
		fmt.Println(k, v)
	}
	//	查看元素是否存在
	name, ok := siteMap["facebook"]
	if ok {
		fmt.Println("facebook 存在", name)
	} else {
		fmt.Println("facebook 不存在")
	}
}
