package main

import (
	option "08"
	"fmt"
)

func main() {
	fmt.Println(option.Lines())
}
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}
