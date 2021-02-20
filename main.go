package main

import (
	"fmt"
	"github.com/jorgealves/go-file-diff/cmd"
	"io"
	"os"
)

func main() {
	fileA, fileB := os.Args[1], os.Args[2]
	result := cmd.CompareJsonFiles(fileA, fileB)
	var output string
	if result {
		output = fmt.Sprintf("Source file: %s\nTarget file: %s\nresult : Files match 👍!\n", fileA, fileB)
	} else {
		output = fmt.Sprintf("Source file: %s\nTarget file: %s\nresult : Files do not match 👎!\n", fileA, fileB)
	}
	_, _ = io.WriteString(os.Stdout, output)
}
