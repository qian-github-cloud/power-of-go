package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type count struct {
	Input     io.Reader
	Output    io.Writer
	wordCount bool
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
func Lines() int {

	c, err := NewCount(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Lines()
}

func WithInputFromArgs(args []string) option {
	return func(c *count) error {
		if len(args) == 0 {
			return nil
		}

		f, err := os.Open(args[0])

		if err != nil {
			return err
		}
		c.Input = f
		return nil
	}
}

func Words() int {
	c, err := NewCount(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return c.Words()
}

func (c *count) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.Input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return words
}

func FromArgs(args []string) option {
	return func(c *count) error {
		fmt.Println("os-Args", os.Args[0])
		//fmt.Println(args[0])
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

		wordCount := fset.Bool("w", false, "Count words instead of lines")

		err := fset.Parse(args)

		if err != nil {
			return err
		}

		c.wordCount = *wordCount

		args := fset.Args()
		fmt.Println("args", args)
		if len(args) < 1 {
			return nil
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.Input = f
		return nil
	}
}

func RunCLi() {
	c, err := NewCount(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Println(c.Words())
	} else {
		fmt.Println(c.Lines())
	}
}
