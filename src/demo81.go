package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		// 一换行符进行分割
		// 读取该段数据 并打印
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
	return
}
func main() {
	// 开始解析参数
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Println("读取出错了")
			continue
		}

		//cat(bufio.NewReader(f))
		// 使用切片来读取文件
		catF(f)
		f.Close()
	}
}

// 使用切片来读取文件
func catF(f *os.File) {
	fmt.Println("使用切片来读取")
	const nbuf = 512
	var buf [nbuf]byte
	for {
		nr, err := f.Read(buf[:])
		if err != nil {
			break
		}
		switch true {
		case nr < 0:
			fmt.Println("小于零")
			os.Exit(1)
		case nr == 0:
			fmt.Println("等于0")
			return
		case nr > 0:
			// 这是直接写入到控制台显示
			nw, _ := os.Stdout.Write(buf[0:nr])
			if nw != nr {
				fmt.Println("最终问题")
			}
			//fmt.Println(nw, nw, nr, buf[0:nr])

		}
	}
}
