package formats

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testdataPath() string {
	path, err := filepath.Abs("../testdata/")

	if err != nil {
		panic("Can't find test file")
	}

	return path
}

func TestCreateMedia_Jpg(t *testing.T) {

	file := "cat.jpg"

	actual, err := getRecord(file)

	if err != nil {
		t.Error(err)
	}

	expected := ".jpg\tcat\t" + testdataPath() + "\t10820\t---\t474\t267\t474x267\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".gif\tcat\t" + testdataPath() + "\t770475\t---\t500\t311\t500x311\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".png	cat	" + testdataPath() + "	62873	---	235	172	235x172	N/A	2025-09-28T15:54:26	---"

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

	expected := ".cr2\tsample1\t" + testdataPath() + "\t67127952\t---\t8688\t5792\t8688x5792\tN/A\t2025-09-28T15:54:27\t---"

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

	expected := ".dng\tsample1\t" + testdataPath() + "\t6372698\t---\t256\t171\t256x171\tN/A\t2025-09-28T15:54:27\t---"

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

	expected := ".mov\tsample\t" + testdataPath() + "\t709764\t00:00:31\t480\t270\t480x270\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".avi\tsample\t" + testdataPath() + "\t675840\t00:00:06\t256\t240\t256x240\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".mp4\tsample\t" + testdataPath() + "\t1570024\t00:00:31\t480\t270\t480x270\tN/A\t2025-09-28T15:54:26\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func BenchmarkGetRecordForMp4WithMp4MediaFile(b *testing.B) {

	filePath := filepath.Join("../testdata", "sample.mp4")
	info, _ := os.Stat(filePath)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewMp4MediaFile(filePath, info).GetRecord()
	}
}

func BenchmarkGetRecordForMp4WithMovieMediaFile(b *testing.B) {

	filePath := filepath.Join("../testdata", "sample.mp4")
	info, _ := os.Stat(filePath)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewMovieMediaFile(filePath, info).GetRecord()
	}
}

func TestCreateMedia_Mpg(t *testing.T) {

	actual, err := getRecord("sample.mpg")

	if err != nil {
		t.Error(err)
	}

	expected := ".mpg\tsample\t" + testdataPath() + "\t548754\t00:00:00\t640\t360\t640x360\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".mkv\tsample\t" + testdataPath() + "\t573066\t00:00:13\t640\t360\t640x360\tN/A\t2025-09-28T15:54:26\t---"

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

	expected := ".mp3\tsample\t" + testdataPath() + "\t764176\t00:00:27\t---\t---\t---\tN/A\t2025-09-28T15:54:26\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}
}

func TestCreateNonMedia_Txt(t *testing.T) {
	actual, err := getRecord("åäö.txt")

	if err != nil {
		t.Error(err)
	}

	expected := ".txt\tåäö\t" + testdataPath() + "\t447\t---\t---\t---\t---\tN/A\t2025-09-28T15:54:27\t---"

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
