package file

import (
	"fmt"
	"testing"
)

func TestGetCurrentFiles(t *testing.T) {
	files, err := GetCodeFiles()
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		fmt.Printf("name: %s, suffixç: %s, stem: %s, content: %s\n", file.Name, file.Suffix, file.Stem, file.Content)
	}
}
