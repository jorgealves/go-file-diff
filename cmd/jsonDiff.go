package cmd

import "github.com/jorgealves/go-file-diff/internal"

// CompareJsonFiles takes the source and target filepath and compare its content
func CompareJsonFiles(source string, target string) bool {
	fileAOutput := internal.ParseFile(source)
	fileBOutput := internal.ParseFile(target)

	return internal.CompareOutput(fileAOutput, fileBOutput)
}