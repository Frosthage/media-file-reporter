package formats

import (
	"context"
	"fmt"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"os"
	"path/filepath"
	"time"
)

type AudioMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media AudioMediaFile) GetRecord() (string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, media.path)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	duration := data.Format.Duration()

	ext := filepath.Ext(media.path)
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		ext,
		media.fileInfo.Name()[0:len(media.fileInfo.Name())-len(ext)],
		media.path,
		media.fileInfo.Size(),
		duration, // duration
		"---",    // width
		"---",    // height
		"---",    // width * height
		"---",
	), nil
}

func (media AudioMediaFile) GetPath() string {
	return media.path
}

func NewAudioMediaFile(path string, info os.FileInfo) AudioMediaFile {
	return AudioMediaFile{path: path, fileInfo: info}
}
