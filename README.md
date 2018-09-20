# Gopherhole # 
[![Go Report Card](https://goreportcard.com/badge/github.com/PenguinGovernor/Gopherhole)](https://goreportcard.com/report/github.com/PenguinGovernor/Gopherhole)

Gopherhole is a set of cli tools that encodes data in order to compress it. Its primary form of communication
between tools is Google's protocol buffers. This project is meant only as a learning excercise.

### Download and Install 
1. Install and configure Go for you operating system. [See Here](https://golang.org/doc/install)
2. Once Go is installed and configured, run this command to install the `gopherhole` subtools 

```shell
go get github.com/penguingovernor/gopherhole/...
``` 
3. Finally, install `gopherhole` by symlinking `gopherhole` to anywhere in your `$PATH` variable

Example:

```shell
sudo ln -s $GOPATH/src/github.com/penguingovernor/gopherhole/gopherhole /opt/bin
```

### Documentation 
After installing, you can use `go doc` to get documenation:
```shell
go doc github.com/penguingovernor/gopherhole/cmd/cencode
go doc github.com/penguingovernor/gopherhole/cmd/scompress
go doc github.com/penguingovernor/gopherhole/cmd/cdecode
```

### Running gopherhole

`gopherhole` can be run in the following format:

```shell
gopherhole $INPUTFILE $COMPRESSIONTYPE
```

Example:
```shell
echo "Hello, world\!" > input.txt # Creates input.txt
gopherhole input.txt zip # Creates input.zip
```

`gopherhole` supports the following compression types:

* "gz" for gzip compression
* "tar" for tar archiving
* "zip" for zip archiving

