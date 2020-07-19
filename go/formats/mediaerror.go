package formats

import (
	"fmt"
	"os"
	"path/filepath"
)

type ErrorMediaFile struct {
	fileInfo os.FileInfo
	path     string
	message  string
}

func (media ErrorMediaFile) GetRecord() (string, error) {

	ext := filepath.Ext(media.path)

	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		ext,
		media.fileInfo.Name()[0:len(media.fileInfo.Name())-len(ext)],
		media.path,
		media.fileInfo.Size(),
		"---", // duration
		"---", // width
		"---", // height
		"---", // width * height
		media.message,
	), nil
}

func (media ErrorMediaFile) Error() string {
	return media.message
}

func NewErrorMediaFile(path string, info os.FileInfo, msg string) ErrorMediaFile {
	return ErrorMediaFile{path: path, fileInfo: info, message: msg}
}
