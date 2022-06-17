package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintHelloMessage(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}

	hello.Print()

	//print.PrintTo()

	want := "Hello, world\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("Want %q, but got %q", want, got)
	}
}
