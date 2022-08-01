package main

import (
	"os"
)

func main() {
	cmd := cmd.NewDefau
	cmd.Stdout = os.Stdout
	cmd.Run()
}
