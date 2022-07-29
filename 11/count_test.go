package count_test

import (
	"archive/zip"
	count "filesystem"
	"os"
	"testing"
	"testing/fstest"
)

// func TestFiles(t *testing.T) {
// 	t.Parallel()

// 	want := 5
// 	got := count.Files("testdata/findgo")
// 	if want != got {
// 		t.Errorf("want %d , but got %d", want, got)
// 	}
// }

// you can use the a key-value map instead of accessing the disk
// because the filesystem like a key-value database
func TestFilesInMemory(t *testing.T) {
	t.Parallel()

	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder/test.go":      {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}

	want := 5
	got := count.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestFileOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/findgo")

	want := 5
	got := count.Files(fsys)

	if want != got {
		t.Errorf("want %d ,but got %d", want, got)
	}
}

func BenchmarkFileOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/findgo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder/test.go":      {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count.Files(fsys)
	}
}

func TestFileInZIP(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader(
		"testdata/findgo.zip",
	)

	if err != nil {
		t.Fatal(err)
	}

	want := 5
	got := count.Files(fsys)
	if want != got {
		t.Errorf("want %d , but not got %d", want, got)
	}
}
