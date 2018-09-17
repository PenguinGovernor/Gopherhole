// Package main is a cli server tool that handles protocol
// buffers and compresses the data, then to reencode them
package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
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
	inputPath := flag.String("input", "", "Input file")
	flag.Parse()
	if *inputPath == "" {
		log.Fatalln("missing input arguement")
	}
	// Get the uncomprsesed data and report any errors.
	uncompressedData, err := getUncompressedData(*inputPath)
	if err != nil {
		log.Fatalln(err)
	}
	// Generate Compressed data and report errors.
	compressedDataBytes, err := compressData(uncompressedData)
	if err != nil {
		log.Fatalln(err)
	}
	// Write the output to another pb file.
	if err := writeData(compressedDataBytes); err != nil {
		log.Fatalln(err)
	}
}

func writeData(data *GopherHoleProtocol.Data) error {
	// Marshal the data
	bb, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	// Write the data to the file
	fileName := strings.TrimSuffix(data.FileName, filepath.Ext(data.FileName)) + ".compressed.pb"
	if err := ioutil.WriteFile(fileName, bb, 0644); err != nil {
		return err
	}
	return nil
}

func getUncompressedData(path string) (*GopherHoleProtocol.Data, error) {
	// Get file bytes and report any errors.
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshal the data and report any errors.
	var data GopherHoleProtocol.Data
	if err := proto.Unmarshal(bb, &data); err != nil {
		return nil, err
	}

	// Return the data
	return &data, nil
}

func compressData(uncomprsesed *GopherHoleProtocol.Data) (*GopherHoleProtocol.Data, error) {
	// Write initial headers.
	compressedData := &GopherHoleProtocol.Data{
		IsCompressed:    true,
		FileName:        uncomprsesed.FileName,
		CompressionType: uncomprsesed.CompressionType,
	}

	// Determine how to compress.
	switch uncomprsesed.CompressionType {
	case "zip":
		if err := zipUp(uncomprsesed, compressedData); err != nil {
			return nil, err
		}
		return compressedData, nil
	case "tar":
		if err := tarUp(uncomprsesed, compressedData); err != nil {
			return nil, err
		}
		return compressedData, nil
	case "gz":
		if err := gzipUp(uncomprsesed, compressedData); err != nil {
			return nil, err
		}
		return compressedData, nil

	}

	// If we reached here no comopression type is recognized
	return nil, fmt.Errorf("unknown compression type: %s", uncomprsesed.CompressionType)
}

func zipUp(uncomprsesed, compressed *GopherHoleProtocol.Data) error {
	// Create a buffer to write to.
	buf := new(bytes.Buffer)
	// Create a zip archive
	w := zip.NewWriter(buf)

	// Create the file inside the archive.
	f, err := w.Create(uncomprsesed.FileName)
	if err != nil {
		return err
	}

	// Write the uncompressed bytes to the file.
	if _, err := f.Write(uncomprsesed.Payload); err != nil {
		return err
	}

	// Close the zip archive
	if err := w.Close(); err != nil {
		return err
	}

	// Write the contents of the buffer to the compmressed Data Structure.
	compressed.Payload = buf.Bytes()

	return nil
}

func tarUp(uncomprsesed, compressed *GopherHoleProtocol.Data) error {
	// Create a buffer to write to.
	buf := new(bytes.Buffer)
	// Create a zip archive
	w := tar.NewWriter(buf)

	// Write the headers.
	header := &tar.Header{
		Name: uncomprsesed.FileName,
		Mode: 0666,
		Size: int64(len(uncomprsesed.Payload)),
	}

	if err := w.WriteHeader(header); err != nil {
		return err
	}

	// Write the uncompressed bytes to the file.
	if _, err := w.Write(uncomprsesed.Payload); err != nil {
		return err
	}

	// Close the zip archive
	if err := w.Close(); err != nil {
		return err
	}

	// Write the contents of the buffer to the compmressed Data Structure.
	compressed.Payload = buf.Bytes()

	return nil
}

func gzipUp(uncomprsesed, compressed *GopherHoleProtocol.Data) error {
	buf := new(bytes.Buffer)
	// Create a zip archive
	w := gzip.NewWriter(buf)

	// Write the uncompressed bytes to the file.
	if _, err := w.Write(uncomprsesed.Payload); err != nil {
		return err
	}

	// Close the zip archive
	if err := w.Close(); err != nil {
		return err
	}

	compressed.Payload = buf.Bytes()

	return nil
}
