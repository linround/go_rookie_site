package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var mu sync.Mutex // 这是一个互斥锁
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	// 然而在并发情况下，假如真的有两个请求同一时刻去更新count，
	//那么这个值可能并不会被正确地增加；这个程序可能会引发一个严重的bug：竞态条件。
	//为了避免这个问题，我们必须保证每次修改变量的最多只能有一个goroutine，
	//这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的
	mu.Lock()
	count++
	mu.Unlock()
	lissajous(w)
}

func lissajous(out io.Writer) {
	const (
		whiteIndex = iota
		blackIndex
	)
	var palette = []color.Color{color.White, color.Black}

	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.00; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d \n", count)
	mu.Unlock()
}
