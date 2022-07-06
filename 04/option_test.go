package option_test

import (
	"bytes"
	"option"
	"testing"
)

func TestOption(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := option.NewCount(
		option.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 3

	got := c.Lines()

	if want != got {
		t.Errorf("want %d , But got %d", want, got)
	}
}
