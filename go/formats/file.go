package formats

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
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

func timeToString(fileTime syscall.Filetime) string {

	t := time.Unix(0, fileTime.Nanoseconds())

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func getLastWriteTime(fileInfo os.FileInfo) string {
	if data, ok := fileInfo.Sys().(*syscall.Win32FileAttributeData); ok {
		data.CreationTime.Nanoseconds()

		return timeToString(data.LastWriteTime)
	}

	return "---"
}

func getCreationTime(fileInfo os.FileInfo) string {

	if data, ok := fileInfo.Sys().(*syscall.Win32FileAttributeData); ok {
		data.CreationTime.Nanoseconds()

		return timeToString(data.CreationTime)
	}

	return "---"
}
