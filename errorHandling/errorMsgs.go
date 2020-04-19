package errorHandling

import (
	"fmt"
)

var MissingCommandErr = fmt.Errorf("A command must be specified to run (store, retrieve, defragment, display)")
var IncorrectNumArgsForStoreErr = fmt.Errorf("The store comamnd requires 2-3 input args with it (<key> <value> <OPTIONAL - databaseFilePath>)")
var IncorrectNumArgsForRetrieveErr = fmt.Errorf("The store comamnd requires 1-2 input args with it (<key> <OPTIONAL - databaseFilePath>)")
var IncorrectNumArgsForDefragmentErr = fmt.Errorf("The defragment comamnd requires 0-1 input args with it (<OPTIONAL - databaseFilePath>)")
var IncorrectNumArgsForDisplayErr = fmt.Errorf("The defragment comamnd requires 0-1 input args with it (<OPTIONAL - databaseFilePath>)")
