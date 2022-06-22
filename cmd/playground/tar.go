package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var buf bytes.Buffer
	err := compress("prometheus/data/01G65HN0X6B0MFQ1GF80RVRFCM", &buf)
	if err != nil {
		panic(err)
	}

	// write the .tar.gzip
	fileToWrite, err := os.OpenFile("./compress.tar.gz", os.O_CREATE|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		panic(err)
	}
}

func compress(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}
