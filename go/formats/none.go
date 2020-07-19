package formats

import (
	"fmt"
	"os"
	"path/filepath"
)

type NonMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (file NonMediaFile) GetRecord() (string, error) {

	ext := filepath.Ext(file.path)

	file.fileInfo.Name()

	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		ext,
		file.fileInfo.Name()[0:len(file.fileInfo.Name())-len(ext)],
		file.path,
		file.fileInfo.Size(),
		"---", // duration
		"---", // width
		"---", // height
		"---", // width * height
		"---",
	), nil
}

func NewNonMediaFile(path string, info os.FileInfo) NonMediaFile {
	return NonMediaFile{path: path, fileInfo: info}
}
