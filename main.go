package main

import (
	"os"

	"database-file-controller/pkg"
	"database-file-controller/errorHandling"
)

func main() {
	argsWithoutProg := os.Args[1:]
	err := pkg.HandleArgs(argsWithoutProg)

	if err != nil {
		errorHandling.HandleError(err)
	}
}
