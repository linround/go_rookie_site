[go语言圣经](https://golang-china.github.io/gopl-zh/ch1/ch1.html)

1. gofmt 自动格式化代码
2. goimports 自动化地添加或删除import声明
3. 关于os.Args，参数从命令后算起：

代码：
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "  "
	}
	fmt.Println(s)
}

```

执行：
> go run a.go -p lomnm -l dsfs dasd   

以上输出如下：
> -p  lomnm  -l  dsfs  dasd

4.运行程序
> 如果你的操作系统是Mac OS X或者Linux，那么在运行命令的末尾加上一个&符号，即可让程序简单地跑在后台，windows下可以在另外一个命令行窗口去运行这个程序。
> 
5. go语言的垃圾回收机制：引用计数