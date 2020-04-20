package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"

	"database-file-controller/globalVars"
)

var readDB = func() error {
	raw_appSpec, err := ioutil.ReadFile(globalVars.DbFilePath)
	globalVars.DbBytes = raw_appSpec

	return err
}

var checkFilePathExists = func() error {
	_, err := os.Stat(globalVars.DbFilePath)

	return err
}

var createDBFile = func() error {
	// Create file
	var file, err = os.Create(globalVars.DbFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("File Successfully Created at", globalVars.DbFilePath)

	// Write Header to new file
	blankHeader := "test"

	sizeOfStr := unsafe.Sizeof(blankHeader)

	fmt.Println(sizeOfStr)
	fmt.Println([]byte(blankHeader))

	_, writeErr := file.Write([]byte(blankHeader))
	if writeErr != nil {
		return writeErr
	}

	dat, _ := ioutil.ReadFile(globalVars.DbFilePath)
    fmt.Println(dat)
    fmt.Println(string(dat))

	// 32-byte checksum

	fmt.Println("Header Successfully Written to", globalVars.DbFilePath)

	return nil
}
