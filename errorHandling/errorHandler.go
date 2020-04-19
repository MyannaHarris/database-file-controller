package errorHandling

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		defer func() {
			fmt.Println("\nERROR: " + err.Error())
			os.Exit(1)
		}()
		panic(err)
	}
}
