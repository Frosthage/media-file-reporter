package formats

import (
	_ "github.com/mdouchement/dng"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

type Media interface {
	GetRecord() ([]string, error)
	GetPath() string
	GetFileInfo() os.FileInfo
}

type MediaRecord []string

func CreateMedia(filePath string, info os.FileInfo) Media {
	switch ext := filepath.Ext(filePath); strings.ToLower(ext) {
	case ".jpg":
		fallthrough
	case ".png":
		fallthrough
	case ".gif":
		fallthrough
	case ".dng":
		fallthrough
	case ".cr2":
		return NewImageMediaFile(filePath, info)
	case ".avi":
		fallthrough
	case ".mpg":
		fallthrough
	case ".mkv":
		return NewMovieMediaFile(filePath, info)
	case ".mov":
		fallthrough
	case ".mp4":
		return NewMp4MediaFile(filePath, info)
	case ".mp3":
		return NewAudioMediaFile(filePath, info)

	default:
	}

	return NewNonMediaFile(filePath, info)
}
