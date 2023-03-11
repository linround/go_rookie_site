package main

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Name      string
	FansCount int64
}

func main() {
	TestJsonUnmarshal()
}

// 反json化， jsonStr是json化的字符串，v传空的结构体

func JsonUnmarshal(jsonStr string, v interface{}) error {
	var fasterJson = jsoniter.ConfigCompatibleWithStandardLibrary
	byteValue := []byte(jsonStr)
	err := fasterJson.Unmarshal(byteValue, v)
	if err != nil {
		return errors.New("JsonUnmarshal failed for v: %+v")
	}
	return nil
}

func TestJsonUnmarshal() {
	const jsonStream = `
        {"name":"ethancai", "fansCount": 9223372036854775807}
    `
	var user User // 类型为User
	err := JsonUnmarshal(jsonStream, &user)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v \n", user)
}
