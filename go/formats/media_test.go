package formats

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateMedia_Jpg(t *testing.T) {

	actual, err := getRecord("cat.jpg")

	if err != nil {
		t.Error(err)
	}

	expected := ".jpg\tcat\t../testdata/cat.jpg\t10820\t---\t474\t267\t474x267\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Gif(t *testing.T) {

	actual, err := getRecord("cat.gif")

	if err != nil {
		t.Error(err)
	}

	expected := ".gif\tcat\t../testdata/cat.gif\t770475\t---\t500\t311\t500x311\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Png(t *testing.T) {

	actual, err := getRecord("cat.png")

	if err != nil {
		t.Error(err)
	}

	expected := ".png\tcat\t../testdata/cat.png\t62873\t---\t235\t172\t235x172\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Cr2(t *testing.T) {

	actual, err := getRecord("sample1.cr2")

	if err != nil {
		t.Error(err)
	}

	expected := ".cr2\tsample1\t../testdata/sample1.cr2\t67127952\t---\t8688\t5792\t8688x5792\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Dng(t *testing.T) {

	actual, err := getRecord("sample1.dng")

	if err != nil {
		t.Error(err)
	}

	expected := ".dng\tsample1\t../testdata/sample1.dng\t6372698\t---\t256\t171\t256x171\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Mov(t *testing.T) {

	actual, err := getRecord("sample.mov")

	if err != nil {
		t.Error(err)
	}

	expected := ".mov\tsample\t../testdata/sample.mov\t709764\t30.571s\t480\t270\t480x270\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Avi(t *testing.T) {

	actual, err := getRecord("sample.avi")

	if err != nil {
		t.Error(err)
	}

	expected := ".avi\tsample\t../testdata/sample.avi\t675840\t6.066667s\t256\t240\t256x240\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Mp4(t *testing.T) {

	actual, err := getRecord("sample.mp4")

	if err != nil {
		t.Error(err)
	}

	expected := ".mp4\tsample\t../testdata/sample.mp4\t1570024\t30.526667s\t480\t270\t480x270\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_ErrorTest(t *testing.T) {

	actual, err := getRecordAbsolutPath ("E:\\slask-4-realz\\2017-09-11-Frosthage-0236231.MP4")

	if err != nil {
		t.Error(err)
	}

	expected := ".mp4\tsample\t../testdata/sample.mp4\t1570024\t30.526667s\t480\t270\t480x270\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}




func TestCreateMedia_Mpg(t *testing.T) {

	actual, err := getRecord("sample.mpg")

	if err != nil {
		t.Error(err)
	}

	expected := ".mpg\tsample\t../testdata/sample.mpg\t548754\t0s\t640\t360\t640x360\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Mkv(t *testing.T) {

	actual, err := getRecord("sample.mkv")

	if err != nil {
		t.Error(err)
	}

	expected := ".mkv\tsample\t../testdata/sample.mkv\t573066\t13.346s\t640\t360\t640x360\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateMedia_Mp3(t *testing.T) {

	actual, err := getRecord("sample.mp3")

	if err != nil {
		t.Error(err)
	}

	expected := ".mp3\tsample\t../testdata/sample.mp3\t764176\t27.252s\t---\t---\t---\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}


func getRecordAbsolutPath(fileName string) (string, error) {

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
		return "", err
	}

	record, err := CreateMedia(fileName, fileInfo).GetRecord()

	if err != nil {
		return "", err
	}

	return strings.Join(record, "\t"), nil
}




func getRecord(fileName string) (string, error) {

	filePath := filepath.Join("../testdata", fileName)
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
		return "", err
	}

	record, err := CreateMedia(filePath, fileInfo).GetRecord()

	if err != nil {
		return "", err
	}

	return strings.Join(record, "\t"), nil
}
