package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Post("http://127.0.0.1:8888/test2", "text/json", nil)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("数据：%q", string(data))
}
func checkError(err error) {
	if err != nil {
		log.Fatal("出现错误", err)
	}
}
