package pipeline

import (
	"bytes"

	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStdout(t *testing.T) {
	t.Parallel()
	want := "Hello, World\n"
	p := FromOfString(want)
	buf := &bytes.Buffer{}

	p.Output = buf
	p.Stdout()

	if p.Error != nil {
		t.Fatal(p.Error)
	}

	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
