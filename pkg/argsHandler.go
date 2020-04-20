package pkg

import (
	"fmt"

	"database-file-controller/errorHandling"
	"database-file-controller/globalVars"
)

func HandleArgs(argsLst []string) error {
	lenArgs := len(argsLst)

	if lenArgs < 1 {
		return errorHandling.MissingCommandErr
	}

	switch argsLst[0] {
	case "store":
		fmt.Println("Calling store")
		if lenArgs > 4 || lenArgs < 3 {
			return errorHandling.IncorrectNumArgsForStoreErr
		}

		if lenArgs == 4 {
			globalVars.DbFilePath = argsLst[3]
		}

		return storeData(argsLst[1], argsLst[2])

	case "retrieve":
		fmt.Println("Calling retrieve")
		if lenArgs > 3 || lenArgs < 2 {
			return errorHandling.IncorrectNumArgsForRetrieveErr
		}

		if lenArgs == 3 {
			globalVars.DbFilePath = argsLst[2]
		}

		return retrieveData(argsLst[1])

	case "free":
		fmt.Println("Calling free")
		if lenArgs > 3 || lenArgs < 2 {
			return errorHandling.IncorrectNumArgsForRetrieveErr
		}

		if lenArgs == 3 {
			globalVars.DbFilePath = argsLst[2]
		}

		return freeData(argsLst[1])

	case "defragment":
		fmt.Println("Calling defragment")
		if lenArgs > 2 {
			return errorHandling.IncorrectNumArgsForDefragmentErr
		}

		if lenArgs == 2 {
			globalVars.DbFilePath = argsLst[1]
		}

		return defragmentData()

	case "display":
		fmt.Println("Calling display")
		if lenArgs > 2 {
			return errorHandling.IncorrectNumArgsForDisplayErr
		}

		if lenArgs == 2 {
			globalVars.DbFilePath = argsLst[1]
		}

		return displayData()

	default:
		return errorHandling.MissingCommandErr
	}

	return nil
}
