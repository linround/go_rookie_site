package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		fmt.Println("**********", string(r))
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			fmt.Errorf("read failed:%v", err)
		}
		// ...use r…
	}
}
