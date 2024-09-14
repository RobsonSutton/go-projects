package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

var err error
var basePath = "git"

func main() {
	var filePath = "./README.md"
	var objectType = "blob"

	// Copy basic init file structure from templates folder
	dirFS := os.DirFS("./templates/git")
	err = os.CopyFS(basePath, dirFS)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// create object based on filepath input
	switch objectType {
	case "blob":
		createBlob(filePath)
	default:
		createBlob(filePath)
	}
}

func createBlob(path string) {
	// Read File Contents
	fileContents, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileContents.Close()

	// Hash File Contents
	hash := sha1.New()
	if _, err = io.Copy(hash, fileContents); err != nil {
		log.Fatal(err)
	}
	hashedValue := fmt.Sprintf("%X", hash.Sum(nil))
	fmt.Println(hashedValue)

	// Generate Object path
	folderName := hashedValue[0:2]
	fileName := hashedValue[2:]
	fmt.Println(folderName)
	fmt.Println(fileName)
	err = os.MkdirAll(basePath+"/objects/"+string(folderName), 0750)
	if err != nil {
		log.Fatal(err)
	}

	gzFile, err := os.Create(basePath + "/objects/" + string(folderName) + "/" + string(fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer gzFile.Close()

	// Compress Object - gzip
	gzipWriter := gzip.NewWriter(gzFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, fileContents)
	if err != nil {
		panic(err)
	}

	gzipWriter.Flush()
}

func readBlob(path string) {

}

// Read a blob object

// Create a blob object

// Read a tree object

// Write a tree object

// Create a commit

// Clone a repository
