package count

import (
	"bufio"
	"io"
	"os"
)

type Count struct {
	Input io.Reader
}

func NewCount() *Count {
	return &Count{
		Input: os.Stdin,
	}
}

func (c *Count) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.Input)

	for scanner.Scan() {
		lines++
	}

	return lines
}

func Lines() int {
	return NewCount().Lines()
}
