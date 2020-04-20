package pkg

import (
	"fmt"
	"os"

	"database-file-controller/globalVars"
)

func storeData(key string, value string) error {
	if err := checkFilePathExists(); err != nil && os.IsNotExist(err) {
		if createErr := createDBFile(); createErr != nil {
			return createErr
		}
	} else {
		return err
	}

	return nil
}

func retrieveData(key string) error {
	if err := checkFilePathExists(); err != nil {
		return err
	}

	return nil
}

func freeData(key string) error {
	if err := checkFilePathExists(); err != nil {
		return err
	}

	return nil
}

func defragmentData() error {
	if err := checkFilePathExists(); err != nil {
		return err
	}
	
	return nil
}

func displayData() error {
	if err := checkFilePathExists(); err != nil {
		return err
	}
	
	if err := readDB(); err != nil {
		return err
	}

	fmt.Printf("\nOutput DB file %s:", globalVars.DbFilePath)
	fmt.Println(string(globalVars.DbBytes))
	fmt.Println()

	return nil
}
