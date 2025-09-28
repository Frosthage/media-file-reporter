package formats

import (
	"os"
	"strconv"
)

type ErrorMediaFile struct {
	fileInfo os.FileInfo
	path     string
	message  string
}

func (media ErrorMediaFile) GetRecord() ([]string, error) {

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
		getLastUpdatedTime(media.fileInfo),
		media.message,
	}, nil
}

func (media ErrorMediaFile) Error() string {
	return media.message
}

func (media ErrorMediaFile) GetPath() string {
	return media.path
}

func (media ErrorMediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewErrorMediaFile(path string, info os.FileInfo, msg string) ErrorMediaFile {
	return ErrorMediaFile{path: path, fileInfo: info, message: msg}
}
