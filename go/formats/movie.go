package formats

import (
	"context"
	"fmt"
	"gopkg.in/vansante/go-ffprobe.v2"
	"os"
	"strconv"
	"time"
)

type MovieMediaFile struct {
	fileInfo os.FileInfo
	path     string
}

func (media MovieMediaFile) GetRecord() ([]string, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, media.path)
	if err != nil {
		return []string{}, NewErrorMediaFile(media.path, media.fileInfo, "ffprobe is unable to read file")
	}

	height := data.FirstVideoStream().Height
	width := data.FirstVideoStream().Width
	duration := data.Format.Duration()

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

func (media MovieMediaFile) GetPath() string {
	return media.path
}

func (media MovieMediaFile) GetFileInfo() os.FileInfo {
	return media.fileInfo
}

func NewMovieMediaFile(path string, info os.FileInfo) MovieMediaFile {
	return MovieMediaFile{path: path, fileInfo: info}
}
