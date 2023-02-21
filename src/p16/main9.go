package main

import (
	"fmt"
	"html/template"
)

func main() {
	temp := "<h1>title<h1><div>title</div>"
	str := template.Must(template.New("TName").Parse(temp))
	fmt.Printf("%v", str)
}
