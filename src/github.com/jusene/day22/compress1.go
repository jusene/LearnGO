package main

import (
	"compress/gzip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	gzipFile, err := os.Open("test.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	// 解压缩到一个writer，它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		panic(err)
	}
	defer outfileWriter.Close()

	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		panic(err)
	}
}
