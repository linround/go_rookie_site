package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:",omitempty"`
}

// 解码   字符串解码到结构体中
// 目前  反向解码  数据解码到 map[string]interface{}

func main() {
	p := &Person{
		Name: "dj",
		Age:  18,
	}

	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
