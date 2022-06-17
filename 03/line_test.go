package count_test

import (
	"bytes"
	"count"
	"testing"
)

func TestPrintHelloMessage(t *testing.T) {
	t.Parallel()
	c := count.NewCount()

	c.Input = bytes.NewBufferString("1\n2\n3")
	want := 3

	got := c.Lines()

	if want != got {
		t.Errorf("want %d , but got %d", want, got)
	}

}
