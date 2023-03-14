package main

import "io"

type reader interface {
	read(p []byte) (n int, err error)
}
type closer interface {
	close() error
}

type readeWriter interface {
	reader
	closer
	io.Writer
}
