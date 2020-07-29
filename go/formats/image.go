package formats

import (
	"fmt"
	"image"
	"os"
	"strconv"
)

type ImageMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media ImageMediaFile) GetRecord() ([]string, error) {

	file, err := os.Open(media.path)
	if err != nil {
		return []string{
			getExt(media),
			getNameWithoutExt(media),
			getAbsoluteFolderPath(media),
			strconv.Itoa(int(media.fileInfo.Size())),
			"---", // duration
			"---", // width
			"---", // height
			"---", // width * height
			getCreationTime(media.fileInfo),
			getLastWriteTime(media.fileInfo),
			"Unable to open media",
		}, err

	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return []string{
			getExt(media),
			getNameWithoutExt(media),
			getAbsoluteFolderPath(media),
			strconv.Itoa(int(media.fileInfo.Size())),
			"---", // duration
			"---", // width
			"---", // height
			"---", // width * height
			getCreationTime(media.fileInfo),
			getLastWriteTime(media.fileInfo),
			"Unable to decode image file",
		}, err
	}

	return []string{
		getExt(media),
		getNameWithoutExt(media),
		getAbsoluteFolderPath(media),
		strconv.Itoa(int(media.fileInfo.Size())),
		"---",                                           // duration
		strconv.Itoa(image.Width),                       // width
		strconv.Itoa(image.Height),                      // height
		fmt.Sprintf("%vx%v", image.Width, image.Height), // width * height
		getCreationTime(media.fileInfo),
		getLastWriteTime(media.fileInfo),
		"---",
	}, nil
}

func (media ImageMediaFile) GetPath() string {
	return media.path
}

func (media ImageMediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewImageMediaFile(path string, info os.FileInfo) ImageMediaFile {
	return ImageMediaFile{path: path, fileInfo: info}
}
