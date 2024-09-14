package main

import (
	"crypto/sha1"
	_ "crypto/sha1"
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
	fmt.Printf("% x", hash.Sum(nil))
}

// Read a blob object

// Create a blob object

// Read a tree object

// Write a tree object

// Create a commit

// Clone a repository
