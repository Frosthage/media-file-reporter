package formats

import (
	"os"
	"strconv"
)

type NonMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media NonMediaFile) GetRecord() ([]string, error) {

	media.fileInfo.Name()

	return []string{
		getExt(media),
		getNameWithoutExt(media),
		getAbsoluteFolderPath(media),
		strconv.Itoa(int(media.fileInfo.Size())),
		"---", // duration
		"---", // width
		"---", // height
		"---", // width * height
		getBirthTime(media.fileInfo),
		getLastChangeTime(media.fileInfo),
		"---",
	}, nil

}

func (media NonMediaFile) GetPath() string {
	return media.path
}

func (media NonMediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewNonMediaFile(path string, info os.FileInfo) NonMediaFile {
	return NonMediaFile{path: path, fileInfo: info}
}
