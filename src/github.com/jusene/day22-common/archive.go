package main

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// 创建一个打包文件
	outfile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	// 使用zip包创建函数zipWriter
	zipWriter := zip.NewWriter(outfile)

	// 往打包文件中写入文件
	var filesToArchive = []struct {
		name, body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = fileWriter.Write([]byte(file.body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
