package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
)

type structA struct {
	A string `yaml:"a"`
}
type structB struct {
	B       string `yaml:"b"`
	structA `yaml:",inline"`
}

// ,flow 父结构体无法访问
// , omitempty 父结构体无法访问
// ,inline 父结构体可以访问

var data = `
a: a string from struct A
b: a string from struct B
`

func main() {
	var b structB
	fmt.Println(b.A)
	err := yaml.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(b.structA.A)
	fmt.Println(b.B)
}
