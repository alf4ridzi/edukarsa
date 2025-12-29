package utils

import (
	"os"
	"path/filepath"
)

func CreateFolder() {
	os.MkdirAll(filepath.Join("assets", "images", "submissions"), 0755)
}
