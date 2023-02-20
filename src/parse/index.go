package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type parseError struct {
	index int
	word  string
	err   error
}

func (e *parseError) string() string {
	return fmt.Sprintf("产生的错误：%q", e.word)
}
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}
func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("没有字段")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(parseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
