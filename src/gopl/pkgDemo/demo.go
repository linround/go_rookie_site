package pkgDemo

import "fmt"
import "rookie/src/gopl/pkg2"

var V int = 1000

func main() {
	fmt.Println("pkgDemo main")
}
func init() {
	pkg2.Test()
	fmt.Println("pkgDemo init")
}
func Test() {
	V++
	fmt.Println("pkgDemo test")
}
