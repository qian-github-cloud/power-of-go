package count_test

import (
	count "filesystem"
	"testing"
	"testing/fstest"
)

func TestFiles(t *testing.T) {
	t.Parallel()

	want := 5
	got := count.Files("testdata/findgo")
	if want != got {
		t.Errorf("want %d , but got %d", want, got)
	}
}

// you can use the a key-value map instead of accessing the disk
// because the filesystem like a key-value database
func TestFilesInMemory(t *testing.T) {
	t.Parallel()

	fsys := fstest.MapFS{
		"file.go": {},
		"subfolder/subfolder.go": {},
		"subfolder/test.go": {},
		"subfolder2/another.go" {},
		"subfolder2/file.go" {},
	}

}
