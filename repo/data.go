package repo

import (
	"fmt"
	"io/ioutil"
	"os"
)

const listItemFile = `repo/items.json`

func ReadDataFile() []byte {
	fileContent, err := os.Open(listItemFile)
	if err != nil {
		fmt.Println(`error to get file`)
		return nil
	}
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	return byteResult
}
