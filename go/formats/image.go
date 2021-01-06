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
	defer file.Close()
	if err != nil {
		fmt.Printf("\rKunde inte öppna bildfilen %v\n", media.path)
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
			"Unable to open media",
		}, NewErrorMediaFile(media.path, media.fileInfo, "Kunde inte öppna filen. Ett annat program kan ha låst filen.")

	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Printf("\rKunde inte läsa bildfilen %v\n", media.path)
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
			"Unable to decode image file",
		}, NewErrorMediaFile(media.path, media.fileInfo, "Bildfilen gick ej att läsa.")
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
		getBirthTime(media.fileInfo),
		getLastUpdatedTime(media.fileInfo),
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
