package main

func main() {
	res := sw(2)
	println(res)
}

func sw(i int) int {
	switch i {
	case 0:
		println(0)
		return 0
	case 1:
		println(1)
		return 1
	case 2:
		println(2)
		// 此处return 之后，不会因为 fallthrough的存在，继续执行下一步
		// 可使用return语句来提前结束分支
		//return 2
		fallthrough
	default:
		println("end")
		return 100

	}
}
