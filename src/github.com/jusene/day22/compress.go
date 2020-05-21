package main

import (
	"compress/gzip"
	"os"
)

func main() {
	outFile, err := os.Create("test.gz")
	if err != nil {
		panic(err)
	}

	gzipWriter := gzip.NewWriter(outFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write([]byte("Gopher!\n"))
	if err != nil {
		panic(err)
	}

}
