package pkg

import (
	"testing"

	"database-file-controller/globalVars"
)

func TestHandleArgs_ValidInput(t *testing.T) {
	var tests = []struct {
		name              string
		inputArgs []string
		expectedDBPath string
	}{
		{"store - no db path",
			[]string{"store", "key", "value"},
			globalVars.DEFAULT_DBFILEPATH},
		{"store - with db path",
			[]string{"store", "key", "value", "test.mydb"},
			"test.mydb"},
		{"retrieve - no db path",
			[]string{"retrieve", "key"},
			globalVars.DEFAULT_DBFILEPATH},
		{"retrieve - with db path",
			[]string{"retrieve", "key", "test.mydb"},
			"test.mydb"},
		{"defragment - no db path",
			[]string{"defragment"},
			globalVars.DEFAULT_DBFILEPATH},
		{"defragment - with db path",
			[]string{"defragment", "test.mydb"},
			"test.mydb"},
		{"display - no db path",
			[]string{"display"},
			globalVars.DEFAULT_DBFILEPATH},
		{"display - with db path",
			[]string{"display", "test.mydb"},
			"test.mydb"},
	}

	for _, test := range tests {
		err := HandleArgs(test.inputArgs)
		if err != nil {
			t.Errorf("HandleArgs failed for: %v", test)
		}

		if globalVars.DbFilePath != test.expectedDBPath {
			t.Errorf("HandleArgs failed to set DbFilePath correctly for: %v. Got: %s", test, globalVars.DbFilePath)
		}

		// Reset globalVars.DbFilePath
		globalVars.DbFilePath = globalVars.DEFAULT_DBFILEPATH
	}
}

func TestHandleArgs_InvalidInput(t *testing.T) {
	var tests = []struct {
		name              string
		inputArgs []string
	}{
		{"No command",
			[]string{}},
		{"store - not enough args",
			[]string{"store", "key"}},
		{"store - too many args",
			[]string{"store", "key", "value", "test.mydb", "test"}},
		{"retrieve - not enough args",
			[]string{"retrieve"}},
		{"retrieve - too many args",
			[]string{"retrieve", "key", "value", "test.mydb"}},
		{"defragment - too many args",
			[]string{"defragment", "key", "test.mydb"}},
		{"display - too many args",
			[]string{"display", "key", "test.mydb"}},
	}

	for _, test := range tests {
		err := HandleArgs(test.inputArgs)
		if err == nil {
			t.Errorf("HandleArgs succeeded but should have failed for: %v", test)
		}
	}
}
