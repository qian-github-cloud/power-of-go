package hello

import (
	"fmt"
	"io"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) PrintTo() {
	fmt.Fprint(p.Output, "Hello, world\n")
}
