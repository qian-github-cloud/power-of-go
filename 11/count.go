package count

import (
	"io/fs"
	"os"
	"path/filepath"
)

// func CountGoFiles(folder string, count int) int {

// 	fsys := os.DirFS(folder)

// 	// match, err := fs.Glob(fsys, "*.go")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {

// 		if filepath.Ext(p) == ".go" {
// 			count++
// 		}
// 		return nil
// 	})

// 	for _, f := range files {

// 		if f.IsDir() {
// 			count = CountGoFiles(folder+"/"+f.Name(), count)
// 		}
// 		if path.Ext(f.Name()) == ".go" {
// 			count++
// 		}
// 	}
// 	return count
// }

func Files(path string) (count int) {

	fsys := os.DirFS(path)

	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	return count
}
