package main

import (
	"flag"
	count "flags"
	"fmt"
)

func main() {
	countWords := flag.Bool("w", false, "Count words instead of lines")
	flag.Parse()
	if *countWords {
		fmt.Println("We're counting words!")
		fmt.Println(count.Words())
	}

}
