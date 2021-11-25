package fileops

import (
	"os"
)

type SmartFile struct {
	Name    string
	Path    string
	Ext     string
	NewPath string
	Size    int64
}

type SmartFiles []*SmartFile

func (s *SmartFiles) Len() int {
	return len(*s)
}

func DirExists(dir string) bool {
	if info, err := os.Stat(dir); err == nil || !os.IsNotExist(err) {
		return info.IsDir()
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil || !os.IsNotExist(err)
}
