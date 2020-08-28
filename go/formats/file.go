package formats

import (
	"fmt"
	"gopkg.in/djherbis/times.v1"
	"os"
	"path/filepath"
	"time"
)

func getExt(media Media) string {
	return filepath.Ext(media.GetPath())
}

func getNameWithoutExt(media Media) string {
	ext := getExt(media)
	return media.GetFileInfo().Name()[0 : len(media.GetFileInfo().Name())-len(ext)]
}

func getAbsoluteFolderPath(media Media) string {
	abs, _ := filepath.Abs(media.GetPath())
	return filepath.Dir(abs)
}

func timeToString(t time.Time) string {

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func getLastUpdatedTime(fileInfo os.FileInfo) string {

	timeSpec := times.Get(fileInfo)

	if timeSpec.HasChangeTime() {
		return timeToString(timeSpec.ChangeTime())
	}

	return timeToString(timeSpec.ModTime())
}

func getBirthTime(fileInfo os.FileInfo) string {

	timeSpec := times.Get(fileInfo)

	if timeSpec.HasBirthTime() {
		return timeToString(timeSpec.BirthTime())
	}

	return "N/A"
}
