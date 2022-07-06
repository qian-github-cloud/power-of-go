package main

import "fmt"

func main() {
	e := New("this is  a  test ")
	if e != nil {
		fmt.Println(e)
	}
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
