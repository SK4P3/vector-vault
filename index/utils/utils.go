package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func TimeTrack(start time.Time, name string) {
	fmt.Printf("%s took %.12fs \n", name, time.Since(start).Seconds())
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
