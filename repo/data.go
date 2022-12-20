package repo

import (
	"fmt"
	"io/ioutil"
	"os"
)

const listItemsFile = `repo/items.json`
const usersFile = `repo/users.json`

func readDataFile(f string) []byte {
	fileContent, err := os.Open(f)
	if err != nil {
		fmt.Println(`error to get file`)
		return nil
	}
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	return byteResult
}

func ReadUsersData() []byte {
	return readDataFile(usersFile)
}

func ReadItemsData() []byte {
	return readDataFile(listItemsFile)
}
