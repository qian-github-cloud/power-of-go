package main

import "fmt"

func main() {
	fmt.Println("hello world")
	Plust(2)
}

func Plust(n int) {

	if n > 10 {
		fmt.Println("n > 10")
		n++
	} else {
		fmt.Println("n < 10")
		n--
	}

}
