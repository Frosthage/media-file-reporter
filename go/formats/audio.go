package formats

import (
	"context"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"os"
	"strconv"
	"time"
)

type AudioMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media AudioMediaFile) GetRecord() ([]string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, media.path)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	duration := data.Format.Duration()

	return []string{
		getExt(media),
		getNameWithoutExt(media),
		getAbsoluteFolderPath(media),
		strconv.Itoa(int(media.fileInfo.Size())),
		GetDuration(duration), // duration
		"---",                 // width
		"---",                 // height
		"---",                 // width * height
		getBirthTime(media.fileInfo),
		getLastUpdatedTime(media.fileInfo),
		"---",
	}, nil
}

func (media AudioMediaFile) GetPath() string {
	return media.path
}

func (media AudioMediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewAudioMediaFile(path string, info os.FileInfo) AudioMediaFile {
	return AudioMediaFile{path: path, fileInfo: info}
}
