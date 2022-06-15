package main

import (
	"hello"
	"os"
)

func main() {

	print := &hello.Printer{
		Output: os.Stdout,
	}

	print.PrintTo()
}
