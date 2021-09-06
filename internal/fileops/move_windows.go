package fileops

import (
	"github.com/natefinch/atomic"
)

func Move(src, dest string) error {
	return atomic.ReplaceFile(src, dest)
}
