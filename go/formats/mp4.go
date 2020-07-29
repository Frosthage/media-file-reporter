package formats

import (
	"fmt"
	"github.com/alfg/mp4"
	"github.com/alfg/mp4/atom"
	"os"
	"strconv"
	"time"
)

type Mp4MediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media Mp4MediaFile) GetRecord() ([]string, error) {

	file, _ := mp4.Open(media.path)
	file.Close()

	width, height := getResolution(file)

	duration := time.Duration(file.Moov.Mvhd.Duration / file.Moov.Mvhd.Timescale) * time.Second

	return []string{
		getExt(media),
		getNameWithoutExt(media),
		getAbsoluteFolderPath(media),
		strconv.Itoa(int(media.fileInfo.Size())),
		GetDuration(duration),               // duration
		strconv.Itoa(width),                 // width
		strconv.Itoa(height),                // height
		fmt.Sprintf("%vx%v", width, height), // width * height
		getCreationTime(media.fileInfo),
		getLastWriteTime(media.fileInfo),
		"---",
	}, nil
}

func getResolution(file *atom.File) (width int, height int) {

	for _, trak := range file.Moov.Traks {
		if trak.Tkhd.GetWidth() > 0 && trak.Tkhd.GetHeight() > 0 {
			width = int(trak.Tkhd.GetWidth())
			height = int(trak.Tkhd.GetHeight())
			break
		}
	}

	return width, height
}


func (media Mp4MediaFile) GetPath() string {
	return media.path
}

func (media Mp4MediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewMp4MediaFile(path string, info os.FileInfo) Mp4MediaFile {
	return Mp4MediaFile{path: path, fileInfo: info}
}
