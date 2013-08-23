# HTTPGogo

HTTPGogo is a simple command-line utility that will start a Go static file server in the port and/or path of your choosing.

## Installation

You need Go installed and `GOBIN` in your `PATH`. `GOBIN` is typically the `bin/` directory in `GOPATH`.

```shell
$ go get github.com/ddoolin/httpgogo
$ go install github.com/ddoolin/httpgogo
```

## Usage (CLI)

Usage is simple. HTTPGogo takes two arguments, `-port` and `-path`. Both are optional. If either is not specificied,
the server's port and path will default to 8080 and the current directory, respectively.

For example, `httpgogo -port=3000 -path=~/documents` would server the /home/username/documents/ directory
on localhost:3000.

```shell
$ httpgogo -h
Usage of httpgogo:
  -path="./": The path to serve. Defaults to the current directory.
  -port=8080: The port to run httpgogo on. Defaults to 8080.
```