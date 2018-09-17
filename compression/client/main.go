// Package main is the client CLI compression-request tool to
// generate a protocol buffer for the server-side compression tool.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/penguingovernor/gopherhole/GopherHoleProtocol"
)

func main() {
	// Declare, parse, and verify flags.
	inputPath := flag.String("input", "", "The input file to generate a compression-request for")
	desiredCompression := flag.String("type", "zip", "The desired compression method. Default: zip")
	flag.Parse()
	if *inputPath == "" {
		log.Fatalln("missing input argument")
	}

	// Encode the file, and report any errors.
	encodedData, err := encode(*inputPath, strings.ToLower(*desiredCompression))
	if err != nil {
		log.Fatalf("could not encode %s: %v\n", *inputPath, err)
	}

	// Write the encoded data to a file.
	if err := writeData(encodedData); err != nil {
		log.Fatalf("could not write data to file: %v\n", err)
	}
}

func writeData(data *GopherHoleProtocol.Data) error {
	// Output file we be the name of the file w/o the extension  + .pb .
	outputName := strings.TrimSuffix(data.FileName, filepath.Ext(data.FileName)) + ".pb"
	// Marshall the data, and report any errors.
	outData, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	// Write marshalled data to file, and report any errors
	if err := ioutil.WriteFile(outputName, outData, 0644); err != nil {
		return err
	}
	return nil
}

func encode(path, compression string) (*GopherHoleProtocol.Data, error) {

	// Verify that we have a valid compression type.
	if err := verifyCompression(compression); err != nil {
		return nil, err
	}

	// Extract bytes from file, and report any errors.
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Serialize Information.
	data := &GopherHoleProtocol.Data{
		IsCompressed:    false,
		Payload:         bb,
		FileName:        filepath.Base(path),
		CompressionType: compression,
	}

	// Return the info we serialized, success.
	return data, nil
}

func verifyCompression(compression string) error {
	switch compression {
	case "zip":
		return nil
	case "tar":
		return nil
	case "huffman":
		return nil
	}
	return fmt.Errorf("unknown compression type: %s", compression)
}
