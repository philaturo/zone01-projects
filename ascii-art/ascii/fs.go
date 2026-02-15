package ascii

import (
	"io/fs"
	"os"
)

// LoadBannerFile loads a banner file using the filesystem API
func LoadBannerFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// LoadFromFS allows loading from any fs.FS (used for testing later)
func LoadFromFS(fsys fs.FS, filename string) ([]byte, error) {
	return fs.ReadFile(fsys, filename)
}
