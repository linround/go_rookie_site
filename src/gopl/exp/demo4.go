package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Person struct {
	Name string `mapstructure:"name"`
}

type Friend1 struct {
	Person
}

// mapstructure
// ,squash 与.remain
// ,remain 中，未被映射的值会放在该字段中
// ,squash 中，嵌套的结构被认为是拥有该结构体名字的另一个字段 提升字段
// ,omitempty 在对象解码成字符串是 在只有默认值的情况下不会解码该字段

type Friend2 struct {
	//
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
    { 
      "type": "friend1",
      "person": {
        "name":"dj1"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	var f1 Friend1
	var f2 Friend2
	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
		// 用来进行类型检测
		switch m["type"].(string) {
		case "friend1":
			mapstructure.Decode(m, &f1)

		case "friend2":
			mapstructure.Decode(m, &f2)
		}
	}

	fmt.Println("friend1", f1)
	fmt.Println("friend2", f2)
}
