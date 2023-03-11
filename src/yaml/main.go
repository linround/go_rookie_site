package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type info struct {
	Name string
	Addr string
}

func main() {

	ex, _ := os.Getwd()
	data, _ := ioutil.ReadFile(ex + "/src/yaml/yaml.yaml")
	var i info
	yaml.Unmarshal(data, &i)
	fmt.Println("初始数据", i)
	d, _ := yaml.Marshal(&i)
	fmt.Println("结果：", d)
	var i2 info
	yaml.Unmarshal(d, &i2)
	fmt.Println("再次转化的数据：", i2)

}
