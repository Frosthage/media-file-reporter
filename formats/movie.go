package formats

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
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
		fmt.Printf("\rffprobe kunde inte läsa %v\n", media.path)
		return []string{}, NewErrorMediaFile(media.path, media.fileInfo, "ffprobe kunde inte läsa filen.")
	}

	var height, width int
	var duration time.Duration

	if firstVideoStream := data.FirstVideoStream(); firstVideoStream != nil {
		height = data.FirstVideoStream().Height
		width = data.FirstVideoStream().Width
		duration = data.Format.Duration()
	} else {
		fmt.Printf("%s saknade en video ström", media.path)
	}

	return []string{
		getExt(media),
		getNameWithoutExt(media),
		getAbsoluteFolderPath(media),
		strconv.Itoa(int(media.fileInfo.Size())),
		GetDuration(duration),               // duration
		strconv.Itoa(width),                 // width
		strconv.Itoa(height),                // height
		fmt.Sprintf("%vx%v", width, height), // width * height
		getBirthTime(media.fileInfo),
		getLastUpdatedTime(media.fileInfo),
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
