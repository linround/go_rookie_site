package main

import (
	"gopl.io/ch13/bzip"
	"io"
	"os"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	io.Copy(w, os.Stdin)
	w.Close()
}
