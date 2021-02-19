package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

// Data Model provided by the examples
type Data struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ParseFile reads the file and parses its data into an array of Data
func ParseFile(filepath string) []Data {
	fileABytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(errors.New("error loading file"))
	}
	var output []Data
	err = json.Unmarshal(fileABytes, &output)
	if err != nil {
		fmt.Println(errors.New("failed parsing file: "))
	}

	return output
}


// Exists verify if a given array contains specified provided object
func Exists(obj interface{}, array []Data) (result bool) {
	result = false
	for _, row := range array {
		if row == obj {
			result = true
			break
		}
	}
	return result
}

// CompareOutput takes the output parsed from two files and compares its content
// Note: Doesn't need to be in the same order
func CompareOutput(output []Data, output2 []Data) bool {
	if len(output) != len(output2) {
		return false
	}
	result := true
	for _, obj1 := range output {
		if !Exists(obj1, output2) {
			result = false
			break
		}
	}
	return result
}
