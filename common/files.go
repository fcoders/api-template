package common

import (
	"os"
	"path/filepath"
)

// Exists returns true/false depending if the file indicated in path, exists or not
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetAppPath() string {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return dir
	}

	return ""
}
