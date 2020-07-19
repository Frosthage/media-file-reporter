package formats

import (
	"context"
	"fmt"
	"gopkg.in/vansante/go-ffprobe.v2"
	"os"
	"path/filepath"
	"time"
)

type MovieMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media MovieMediaFile) GetRecord() (string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, media.path)
	if err != nil {
		return "", NewErrorMediaFile(media.path, media.fileInfo, "ffprobe is unable to read file")
	}

	height := data.FirstVideoStream().Height
	width := data.FirstVideoStream().Width
	duration := data.Format.Duration()

	ext := filepath.Ext(media.path)
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v",
		ext,
		media.fileInfo.Name()[0:len(media.fileInfo.Name())-len(ext)],
		media.path,
		media.fileInfo.Size(),
		duration,                            // duration
		width,                               // width
		height,                              // height
		fmt.Sprintf("%vx%v", width, height), // width * height
		"---",
	), nil
}

func NewMovieMediaFile(path string, info os.FileInfo) MovieMediaFile {
	return MovieMediaFile{path: path, fileInfo: info}
}
