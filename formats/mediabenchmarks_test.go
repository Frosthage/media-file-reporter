package formats

import (
	"fmt"
	"testing"
)

func BenchmarkGetRecordMp3(b *testing.B) {

	record, err := getRecord("sample.mp3")

	fmt.Println(record)

	if err != nil {
		fmt.Println("sdfsdfsdfsdf")
		b.Error(err)
	}

	/*expected := ".mp3\tsample\t../testdata/sample.mp3\t764176\t27.252s\t---\t---\t---\t---"

	if actual != expected {
		t.Errorf(
			"\nExpected: '%v' \n"+
				"Got:      '%v'", expected, actual)
	}	*/
}
