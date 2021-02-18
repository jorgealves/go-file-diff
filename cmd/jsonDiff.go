package cmd

import "github.com/jorgealves/go-file-diff/internal"

func CompareJsonFiles(source string, target string) bool {
	fileAOutput := internal.ParseFile(source)
	fileBOutput := internal.ParseFile(target)

	return internal.CompareOutput(fileAOutput, fileBOutput)
}
