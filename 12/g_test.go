package g_test

import (
	"os"
	"testing"

	"gotest.tools/assert/cmp"
)

func TestParseGitStatusOutput(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/status.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := g.Status{
		Message: "Untracked files",
	}

	got, err := g.ParseGitStatusOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}

	// if want := got {
	// 	t.Errorf("want %d but got %d", want,got)
	// }

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
