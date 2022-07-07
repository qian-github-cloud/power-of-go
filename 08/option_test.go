package option_test

import (
	option "08"
	"bytes"
	"testing"
)

func TestOption(t *testing.T) {
	t.Parallel()

	args := []string{"testdata/three_lines.txt"}

	c, err := option.NewCount(
		option.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d , but got is %d", want, got)
	}

}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()

	inputBuf := bytes.NewBufferString("1\n2\n3")

	c, err := option.NewCount(
		option.WithInput(inputBuf),
		option.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d but got %d", want, got)
	}
}
