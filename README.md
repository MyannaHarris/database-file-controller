# database-file-controller
Database backend file conroller - formatter, serializer, deserializer

## Getting Started

You must have Go `1.13` or later installed.

Build the project:

```
$ go build
```

Build and run in one shot:

```
$ go run main.go store <key> <value> <OPTIONAL - databaseFilePath>
$ go run main.go retrieve <key> <OPTIONAL - databaseFilePath>
$ go run main.go defragment <OPTIONAL - databaseFilePath>
$ go run main.go display <OPTIONAL - databaseFilePath>
```

Example commands:
```
$ go run main.go store test-key test-val test.mydb
$ go run main.go retrieve test-key test.mydb
$ go run main.go defragment test.mydb
$ go run main.go display test.mydb
```

Run unit tests

```
$ go test -v ./pkg/*

# Get percentage of test coverage
go test -coverprofile=cover.out ./pkg/*
```

Format code:

```
$ ./gofmt
```

## File Format
```
<HEADER_checksum><HEADER_NumberOfKeys><HEADER_KeyStringSize><HEADER_KeyAddressSize><KEY_1_KeyBytes><KEY_1_ByteAddressOfVal><KEY_2_KeyBytes><KEY_2_ByteAddressOfVal>...
<EMPTY_SPACE?>...
...<VAL_2_checksum><VAL_2_ValSize><VAL_2_data><VAL_1_checksum><VAL_1_ValSize><VAL_1_data>
```

Example:
```
# Keys will be int32

# Values will be strings
# Golang encodes each character with UTF-8
# Golang represents all characters in int32 so 4 bytes?

# Addresses will be the index of the starting byte in a Byte array of the file
# So Addresses will be int64

<aaa...aaa><2_keys><4_bytes><8_bytes><1_key><200_startByteIndex><2_key><400_startByteIndex>
<>
<bbb...bbb><>
```
