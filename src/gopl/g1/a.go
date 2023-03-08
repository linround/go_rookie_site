// 测试
// meth
package main

//text
import (
	"fmt"
	"os"
)

func main() {

	// 卖弄11
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "  "
	}
	fmt.Println(s)
}
