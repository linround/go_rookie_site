package main

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func main() {

	cwd, _ := os.Getwd()
	path := cwd + "\\src\\gopl\\exp8\\img\\img1.jpg"
	//path2 := cwd + "\\src\\gopl\\exp8\\img\\test_p.png"
	var ps = []string{path, path}
	var ch = make(chan string, len(ps))
	for range ch {
		ch <- ps[0]
	}
	makeThumbnails6(ch)

}

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			fmt.Println("err:", err, "\n")
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err == nil {
			return err
		}
	}

	return nil
}

func makeThumbnails3(filenames []string) {
	ch := make(chan string)
	for _, f := range filenames {
		// 创建一个goroutine
		go func(f string) {
			fmt.Println("开始处理")
			thumbnail.ImageFile(f)
			// 在处理完成后发送一个信号
			fmt.Println("处理完成")
			ch <- f
		}(f)
	}
	// Wait for goroutines to complete.
	for range filenames {
		fmt.Println(<-ch)
	}
}

func makeThumbnails(filenames []string) {
	// 迭代图像生成缩略图
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // 利用多核cpu的计算能力来拉伸图像
	}
}

//func makeThumbnails(filenames []string) {
//	// 迭代图像生成缩略图
//	for _, f := range filenames {
//		if _, err := thumbnail.ImageFile(f); err != nil {
//			fmt.Println("出错")
//		}
//	}
//}
