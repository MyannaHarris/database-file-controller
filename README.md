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

```
