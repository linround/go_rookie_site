package main

import (
	"flag"
	"os"
)

// flag 包用来解析命令行选项

//go run demo80.go -dddd=true asdas asd-asdas -sdasda -asda
//-dddd
//新的一行
//nArg: 4
//asdas
//asd-asdas
//-sdasda
//-asda

var NewLine = flag.Bool("d", false, "新的一行")

func main() {
	//flag.PrintDefaults()
	// flag.Parse() 扫描参数列表（或者常量列表）并设置 flag
	flag.Parse()
	var s string = ""
	//nArg := flag.NArg()
	//fmt.Println("nArg:", nArg)
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += "\n"
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
