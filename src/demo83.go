package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type address struct {
	city    string
	country string
}
type card struct {
	address []*address
	remark  string
}

func main() {
	ex, err := os.Getwd()
	if err != nil {
		fmt.Println("出错了")
		return
	}

	jsonFile := ex + "/src/vcard.json"

	pa := &address{"sz", "cn"}
	wa := &address{"gz", "cn"}
	vc := card{[]*address{pa, wa}, "sad"}
	fmt.Println("vc.remark", vc.remark)
	// json 格式化
	js, _ := json.Marshal(vc)
	fmt.Printf("json: %s \n", js)

	// 解码json
	var v interface{}
	json.Unmarshal(js, &v)
	fmt.Println("v.remark \n", v)

	// 使用编码
	file, _ := os.OpenFile(jsonFile, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err = enc.Encode(vc)

	if err != nil {
		log.Println("出错了", err)
	}
}
