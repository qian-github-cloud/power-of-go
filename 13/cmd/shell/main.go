package main

import (
	"bufio"
	"fmt"
	"os"
	"shell"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("\n Be seeing you!")
			break
		}

		cmd, err := shell.CmdFromString(line)
		if err != nil {
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error :", err)
		}

		fmt.Printf("%s", out)
	}

}
