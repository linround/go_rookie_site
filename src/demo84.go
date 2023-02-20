package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("出错了啊")
	}
	fmt.Printf("f: %s \n", f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)

		case float64:
			fmt.Println(k, "is float64", vv)

		case []interface{}:
			fmt.Println(k, "is arr", vv)

		default:
			fmt.Println(k, "未知类型")

		}
	}

	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	// 把json数据写入文件
	jsonFile := ex + "/src/vcard.json"
	file, _ := os.OpenFile(jsonFile, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	// 初始化文件
	enc := json.NewEncoder(file)
	// 开始写入json数据
	err = enc.Encode(m)
}
