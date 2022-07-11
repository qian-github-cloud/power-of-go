package count_test

import (
	"bytes"
	count "flags"
	"testing"
)

func TestWords(t *testing.T) {

	t.Parallel()

	inputBuf := bytes.NewBufferString("1\n2 words\n3 this time")

	c, err := count.NewCount(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 6
	got := c.Words()

	if want != got {
		t.Errorf("want %d , but got %d", want, got)
	}

}

func TestWordsWithFlag(t *testing.T) {
	t.Parallel()

	args := []string{"-w", "testdata/three_lines.txt"}

	c, err := count.NewCount(
		count.FromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}

	want := 6
	got := c.Words()

	if want != got {
		t.Errorf("want %d , but got %d", want, got)
	}

}
