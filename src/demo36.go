package main

import "time"

func main() {
	now := time.Now()
	day := now.Day()
	month := now.Month()
	year := now.Year()

	println(now.Format(time.ANSIC))
	println(now.Format(time.RFC822))
	println(now.Format("02 Jan 2006 15:04"))
	println(day, month, year)
	// 经过一定实际执行
	time.After(500)
	// 经过一点周期执行
	time.Tick(200)
	// 实现对某个进程（实际上是goroutine）暂停某段时长
	time.Sleep(200)
}
