// cdecode is a client cli decoding tool. It takes a protocol
// buffer and decodes it into a properly formatted compressed format.
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/penguingovernor/gopherhole/GopherHoleProtocol"
)

func main() {
	// Define, parse, and verify flags.
	path := flag.String("input", "", "The input")
	flag.Parse()
	if *path == "" {
		log.Fatalln("missing input areguments")
	}
	// Get the data.
	data, err := getData(*path)
	if err != nil {
		log.Fatalln(err)
	}
	// Write the zip file
	writeCompressed(data)
}

func writeCompressed(data *GopherHoleProtocol.Data) error {
	outputName := strings.TrimSuffix(data.FileName, filepath.Ext(data.FileName)) + "." + data.CompressionType
	return ioutil.WriteFile(outputName, data.Payload, 0666)
}

func getData(path string) (*GopherHoleProtocol.Data, error) {
	// Get the bytes.
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshal the bytes.
	var data GopherHoleProtocol.Data
	if err := proto.Unmarshal(bb, &data); err != nil {
		return nil, err
	}
	// Return the data
	return &data, nil
}
