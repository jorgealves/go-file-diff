package main

import (
	"fmt"
	"github.com/jorgealves/go-file-diff/cmd"
	"os"
)

func main() {
	fileA, fileB := os.Args[1], os.Args[2]

	result := cmd.CompareJsonFiles(fileA, fileB)

	fmt.Println(result)

}
