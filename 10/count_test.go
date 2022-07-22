package count_test

import (
	"bytes"
	count "file"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	want := 4
	got := c.Words()

	if want != got {
		t.Errorf("want %d , but got %d", want, got)
	}

}

func TestFromArgsErrorsOnBogusFlag(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}

	_, err := count.NewCount(
		//count.WithOutput(io.Discard),
		count.FromArgs(args),
	)

	if err == nil {
		t.Fatal("want error on bogus flag , got nil")
	}
}

func TestWroteToFile(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/write_test.txt"

	err := os.WriteFile(path, []byte{4, 5, 6}, 0600)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte{1, 2, 3}
	err = count.WriteToFile(path, want)

	if err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

func TestPermsClosed(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/perms_test.txt"

	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = count.WriteToFile(path, []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, but got 0%o", perm)
	}

}
