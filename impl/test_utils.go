package impl

import (
	"io/ioutil"
	"testing"
)

func getTempFile(t *testing.T, pattern string) string {
	tempFile, err := ioutil.TempFile("", pattern)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	filename := tempFile.Name()
	tempFile.Close()
	return filename
}
