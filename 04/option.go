package option

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type count struct {
	Input  io.Reader
	Output io.Writer
}

//
type option func(*count) error

func NewCount(opts ...option) (count, error) {

	c := count{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return count{}, err
		}
	}
	return c, nil
}

//
func WithInput(input io.Reader) option {
	return func(c *count) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.Input = input
		return nil
	}
}

//
func WithOutput(output io.Writer) option {
	return func(c *count) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		c.Output = output
		return nil
	}

}

//
func (c count) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

//
func Lines() (int, error) {
	c, err := NewCount()
	if err != nil {
		return 0, err
	}
	return c.Lines(), nil
}
