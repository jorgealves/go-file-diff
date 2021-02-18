package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"reflect"
)

type Data struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

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


func CompareOutput(output []Data, output2 []Data) bool {
	return reflect.DeepEqual(output, output2)
}