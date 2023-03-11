package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

var document3 = `{
"categories":["12","222","333"],
"people":{
	"name":"jack",
	"age":{
		"birth":10,
		"year":2000,
		"animals":[
				{"barks":"yes","tail":"yes"},
				{"barks":"no","tail":"yes"}
			]
		}
	}
}`

type Items1 struct {
	Categories []string `mapstructure:"categories"`
	//
	Peoples People1 `mapstructure:"People"` // Specify the location of the array
}

type People1 struct {
	// people 大结构传进来
	// 这个结构中的age 对应people结构中 的 age.birth
	Age     int                    `mapstructure:"age.birth"` // jpath is relative to the array
	Animals []Animal               `mapstructure:"age.animals"`
	Other   map[string]interface{} `mapstructure:",remain"`
}
type Animal struct {
	Barks string `jpath:"barks"`
}

func main() {
	TestMetaData()
}

func TestMetaData() {
	var items1 Items1
	// 配置解析规则
	config := &mapstructure.DecoderConfig{
		Metadata: &mapstructure.Metadata{
			Keys:   nil,
			Unused: nil,
		},
		Result: &items1,
	}
	// 根据规则获得解码器
	decoder, _ := mapstructure.NewDecoder(config)
	// 生成二进制数据
	docScript := []byte(document3)
	var docMap map[string]interface{}
	// 解码数据 填入docMap
	_ = json.Unmarshal(docScript, &docMap)

	_ = decoder.Decode(docMap)
	fmt.Println(items1.Peoples.Age)
	fmt.Println(config.Metadata.Keys)
	fmt.Println(config.Metadata.Unused)
}
