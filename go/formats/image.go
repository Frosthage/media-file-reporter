package formats

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
)

type ImageMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media ImageMediaFile) GetRecord() (string, error) {
	ext := filepath.Ext(media.path)

	file, err := os.Open(media.path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", media.path, err)
		return "", err
	}

	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		ext,
		media.fileInfo.Name()[0:len(media.fileInfo.Name())-len(ext)],
		media.path,
		media.fileInfo.Size(),
		"---",        // duration
		image.Width,  // width
		image.Height, // height
		fmt.Sprintf("%vx%v", image.Width, image.Height), // width * height
		"---",
	), nil
}

func (media ImageMediaFile) GetPath() string {
	return media.path
}

func NewImageMediaFile(path string, info os.FileInfo) ImageMediaFile {
	return ImageMediaFile{path: path, fileInfo: info}
}
