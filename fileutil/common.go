package fileutil

import (
	"errors"
	"log"
	"os"
)

func CanOpen(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("file or dir %s does not exist", path)
		}
		log.Fatalf("file or dir %s with unexpected error %v", path, err)
		return false
	}
	return true
}

func IsDir(path string) bool {
	if canOpen := CanOpen(path); !canOpen {
		return false
	}
	fi, _ := os.Stat(path)
	return fi.IsDir()
}

func IsFile(path string) bool {
	if canOpen := CanOpen(path); !canOpen {
		return false
	}
	return !IsDir(path)
}

func FileSize(path string) int64 {
	if isFile := IsFile(path); !isFile {
		return 0
	}
	f, err := os.Stat(path)
	if err != nil {
		log.Fatalf("open %v: %v", path, err)
		return 0
	}
	return f.Size()
}
