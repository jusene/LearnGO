package main

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func main() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	// 遍历打包文件中的每一个文件和文件夹
	for _, file := range zipReader.Reader.File {
		zipFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zipFile.Close()

		targetDir := "./"
		extraFile := filepath.Join(targetDir, file.Name)

		if file.FileInfo().IsDir() {
			log.Println("正在创建目录: ", extraFile)
			os.MkdirAll(extraFile, file.Mode())
		} else {
			log.Println("正在提取文件: ", file.Name)

			outFile, err := os.OpenFile(
				extraFile,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode())
			if err != nil {
				panic(err)
			}

			defer outFile.Close()

			_, err = io.Copy(outFile, zipFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
