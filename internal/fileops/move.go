package fileops

import "os"

func Move(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		return err
	}

	return nil
}
