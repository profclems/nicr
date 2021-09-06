package fileops

import (
	"os"
	"path/filepath"
)

type FileType int

const (
	Applications FileType = iota
	Archive
	Audio
	Documents
	Databases
	Fonts
	Other
	Pictures
	Videos
)

type SmartFile struct {
	Name string
	Path string
	Type FileType
}

type SmartFiles []*SmartFile

func (s *SmartFiles) Len() int {
	return len(*s)
}

func GetFileType(file string) FileType {
	switch filepath.Ext(file) {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".cr2", ".tif", ".bmp", ".heif", ".jxr", ".psd", ".ico", ".dwg":
		return Pictures
	case ".mp4", ".m4v", ".mkv", ".webm", ".mov", ".avi", ".wmv", ".mpg", ".flv", ".3gp":
		return Videos
	case ".wasm", ".dex", ".dey", ".exe", ".dmg", ".rpm", ".deb":
		return Applications
	case ".woff", ".woff2", ".ttf", ".otf":
		return Fonts
	case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".epub", ".rtf", ".txt":
		return Documents
	case ".mid", ".mp3", ".m4a", ".ogg", ".flac", ".wav", ".amr", ".aac":
		return Audio
	case ".zip", ".tar", ".rar", ".gz", ".bz2", ".7z", ".xz", ".zstd", ".swf", ".iso", ".eot", ".ps", ".nes", ".crx", ".cab", ".ar", ".Z", ".lz", ".elf", ".dcm":
		return Archive
	case ".sqlite", ".sql":
		return Databases
	default:
		return Other
	}
}

func DirExists(dir string) bool {
	if info, err := os.Stat(dir); err == nil || !os.IsNotExist(err) {
		return info.IsDir()
	}
	return false
}
