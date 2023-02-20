package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//依据指定的格式向第一个参数内写入字符串，
	//第一参数必须实现了 io.Writer 接口。
	//  Fprintf() 能够写入任何类型，只要其实现了 Write 方法
	fmt.Fprintf(os.Stdout, "%s\n", "hello world!  -  unbuffered")

	//	 buffer
	// bufio.Writer 实现了 Write 方法：
	//它还有一个工厂函数：传给它一个 io.Writer 类型的参数，
	//它会返回一个缓冲的 bufio.Writer 类型的 io.Writer
	buf := bufio.NewWriter(os.Stdout)
	//	添加buffer
	fmt.Fprintf(buf, "%s \n", "hello world! - buffered")
	// 缓冲写入的最后千万不要忘了使用 Flush()
	// 否则最后的输出不会被写入
	buf.Flush()
}
