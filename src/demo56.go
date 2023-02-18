package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	next := Node{data: 200, next: &Node{data: 300}}
	node := Node{data: 100, next: &next}
	myFile := NewFile(100, "myFile")
	fmt.Println(node.next.next)
	fmt.Println(myFile)
	fmt.Println(myFile.fd)
	fmt.Println(myFile.name)
}

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
